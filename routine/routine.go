/**
* @file routine.go
* @brief routine knowledge
* @author ligang
* @version 1.0
* @date 2014-10-10
 */

package routine

const (
	ROUTINE_DONE      = 1
	ROUTINE_START_END = 2
)

type T_Routine_Status chan int
