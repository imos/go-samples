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
  // The -l flag is a line-buffered mode, and it is required for interaction.
  command := exec.Command("sed", "-le", "s/xxx/zzz/")
  raw_stdin, _ := command.StdinPipe()
  stdin := bufio.NewWriter(raw_stdin)
  raw_stdout, _ := command.StdoutPipe()
  stdout := bufio.NewReader(raw_stdout)
  command.Start()
  for i := 0; i < 3; i++ {
    size, _ := stdin.WriteString("aaaxxxccc\n")
    if size != 10 {
      t.Errorf("Output size should be 10, but %#v.", size)
    }
    stdin.Flush()
    actual, _ := stdout.ReadString('\n')
    expected := "aaazzzccc\n"
    if actual != expected {
      t.Errorf("Output should be %#v, but %#v.", expected, actual)
    }
  }
  raw_stdin.Close()
}
