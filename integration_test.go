package netsetgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("netsetgo binary", func() {
	Context("when passed all required args", func() {
		var (
			parentPid int
			pid       int
		)

		BeforeEach(func() {
			cmd := exec.Command("sh", "-c", "ip netns add testNetNamespace")
			Expect(cmd.Run()).To(Succeed())

			cmd = exec.Command("sh", "-c", "ip netns exec testNetNamespace sleep 1000")
			Expect(cmd.Start()).To(Succeed())

			parentPid = cmd.Process.Pid

			cmd = exec.Command("sh", "-c", fmt.Sprintf("ps --ppid %d | tail -n 1 | awk '{print $1}'", parentPid))
			pidBytes, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())

			pid, err = strconv.Atoi(strings.TrimSpace(string(pidBytes)))
			Expect(err).NotTo(HaveOccurred())

			command := exec.Command(pathToNetsetgo,
				"--bridgeName=tower",
				"--bridgeAddress=10.10.10.1/24",
				"--vethNamePrefix=v",
				fmt.Sprintf("--pid=%d", pid),
			)

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		AfterEach(func() {
			parentProcess, err := os.FindProcess(parentPid)
			Expect(err).NotTo(HaveOccurred())

			Expect(parentProcess.Kill()).To(Succeed())

			cmd := exec.Command("sh", "-c", "ip netns delete testNetNamespace")
			Expect(cmd.Run()).To(Succeed())
			cmd = exec.Command("sh", "-c", "ip link delete tower")
			Expect(cmd.Run()).To(Succeed())
			cmd = exec.Command("sh", "-c", "ip link delete v0") // will implicitly delete v1 :D
			Expect(cmd.Run()).To(Succeed())
		})

		It("creates a bridge device on the host with the provided name", func() {
			_, err := net.InterfaceByName("tower")
			Expect(err).NotTo(HaveOccurred())
		})

		It("assignes the provided IP address to the bridge", func() {
			bridge, err := net.InterfaceByName("tower")
			Expect(err).NotTo(HaveOccurred())

			bridgeAddresses, err := bridge.Addrs()
			Expect(err).NotTo(HaveOccurred())

			Expect(bridgeAddresses[0].String()).To(Equal("10.10.10.1/24"))
		})

		// TODO: why does the link go down after a veth is attached?
		PIt("sets the bridge link up", func() {
			Eventually(func() string {
				carrierFileContents, err := ioutil.ReadFile("/sys/class/net/tower/carrier")
				Expect(err).NotTo(HaveOccurred())
				return string(carrierFileContents)
			}).Should(Equal("1\n"))
		})

		It("creates a veth pair on the host using the provided name prefix", func() {
			_, err := net.InterfaceByName("v0")
			Expect(err).NotTo(HaveOccurred())
		})

		It("attaches the host's side of the veth pair to the bridge", func() {
			Expect("/sys/class/net/v0/master").To(BeAnExistingFile())
		})

		It("puts the container's side of the veth pair into the net ns of the process specified by the provided pid", func() {
			stdout := gbytes.NewBuffer()
			cmd := exec.Command("sh", "-c", "ip netns exec testNetNamespace ip addr")
			_, err := gexec.Start(cmd, stdout, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(stdout).Should(gbytes.Say("v1"))
		})
	})
})
