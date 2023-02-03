package dashboard

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func timeSinceModified(filename string) (time.Duration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	modifiedTime := fileInfo.ModTime()
	timeSinceModified := time.Since(modifiedTime)
	return timeSinceModified, nil
}

func isModifiedOver2MinutesAgo(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}

	modifiedTime := fileInfo.ModTime()
	timeSinceModified := time.Since(modifiedTime)
	if timeSinceModified > (2 * time.Minute) {
		return true, nil
	}
	return false, nil
}

func mrefdUptime(pidFile, checkFile string) (time.Duration, error) {
	uptime, err := timeSinceModified(pidFile)
	if err != nil {
		return 0, err
	}
	stale, err := isModifiedOver2MinutesAgo(checkFile)
	if err != nil {
		return 0, err
	}
	if stale {
		return uptime, fmt.Errorf("check file is older than two minutes")
	}
	return uptime, nil
}

func mrefdCheckService() ([]byte, error) {
	cmd := exec.Command("systemctl", "check", "mrefd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
		}
	}
	return out, nil
}

// if needed to test the function
/*func main() {
	pidFile := "/var/run/mrefd.pid"
	checkFile := "/var/log/mrefd.xml"
	uptime, err := mrefdUptime(pidFile, checkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("mrefd has been up for %s", uptime)
}
*/
