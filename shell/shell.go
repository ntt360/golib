/**
* @file shell.go
* @brief shell
* @author ligang
* @version 1.0
* @date 2014-10-10
 */

package shell

import (
	//     "fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

type T_Shell_Result struct {
	Ok     bool
	Output string
}

/**
* @brief run cmd use bash -c
*
* @param string
*
* @return T_Shell_Result
 */
func RunCmd(cmd_str string) *T_Shell_Result {
	result := &T_Shell_Result{
		Ok:     true,
		Output: "",
	}

	cmd := exec.Command("/bin/bash", "-c", cmd_str)
	output, err := cmd.CombinedOutput()
	result.Output = string(output)

	if nil != err {
		result.Ok = false
		result.Output += err.Error()
	}
	return result
}

/**
* @brief run shell as special user
*
* @param string
* @param string
*
* @return
 */
func RunAsUser(cmd_str string, username string) *T_Shell_Result {
	var cmd string

	cur_user, _ := user.Current()
	if "root" == cur_user.Name {
		cmd += "/sbin/runuser " + username + " -c \""
		cmd += strings.Replace(cmd_str, "\"", "\\\"", -1)
		cmd += "\""
	} else {
		cmd += "sudo -u " + username + " "
		cmd += cmd_str
	}

	return RunCmd(cmd)
}

/**
* @brief exec rsync
*
* @param string
* @param string
* @param string
* @param string
* @param string
*
* @return
 */
func Rsync(host string, sou string, dst string, exclude_from string, ssh_user string, timeout int) *T_Shell_Result {
	rsync_cmd := MakeRsyncCmd(host, sou, dst, exclude_from, timeout)

	return RunAsUser(rsync_cmd, ssh_user)
}

func MakeRsyncCmd(host string, sou string, dst string, exclude_from string, timeout int) string {
	to := strconv.Itoa(timeout)
	rsync_cmd := "/usr/bin/rsync -av -e 'ssh -o StrictHostKeyChecking=no -o ConnectTimeout=" + to + "' --timeout=" + to + " --update "
	_, err := os.Stat(exclude_from)
	if nil == err {
		rsync_cmd += "--exclude-from='" + exclude_from + "' "
	}
	rsync_cmd += sou + " " + host + ":" + dst + " 2>&1"

	return rsync_cmd
}
