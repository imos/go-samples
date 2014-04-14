package exec

import (
  "bufio"
  "os/exec"
  "testing"
)

func TestSimpleRead(t *testing.T) {
  command := exec.Command("echo", "hoge", "piyo")
  command_output, _ := command.Output()
  actual := string(command_output)
  expected := "hoge piyo\n"
  if actual != expected {
    t.Errorf("Command line output should be %#v, but %#v.", expected, actual)
  }
}

func TestInteraction(t *testing.T) {
  command := exec.Command("cat")
  raw_stdin, _ := command.StdinPipe()
  stdin := bufio.NewWriter(raw_stdin)
  raw_stdout, _ := command.StdoutPipe()
  stdout := bufio.NewReader(raw_stdout)
  command.Start()
  for i := 0; i < 3; i++ {
    size, _ := stdin.WriteString("abc\n")
    if size != 4 {
      t.Errorf("Output size should be 4, but %#v.", size)
    }
    stdin.Flush()
    actual, _ := stdout.ReadString('\n')
    expected := "abc\n"
    if actual != expected {
      t.Errorf("Output should be %#v, but %#v.", expected, actual)
    }
  }
  raw_stdin.Close()
}
