package main

import (
"fmt"
"os/exec"
"strings"
"iter"
"slices"
)

func main() {
	deforest()
}

func deforest() error {
	op1, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return err;
	}
	curBranch := strings.TrimSpace(string(op1))
	fmt.Println("Branch:", curBranch)
	op2, err := exec.Command("git", "refs", "list", "refs/heads/", "--format", "%(refname:short)").Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return err;
	}
	branches := slices.Collect(Filter(strings.SplitSeq(string(op2), "\n"), func (branch string) bool { return len(branch) > 0 && branch != curBranch }))
	fmt.Printf("Branches: %#v, %d\n", branches, len(branches))
	if len(branches) == 0 {
		fmt.Println("No branches to remove!")
		return fmt.Errorf("No branches to remove")
	}
	args := append([]string {"branch", "-d"}, branches...)
	op3, err := exec.Command("git", args...).Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return err;
	}
	fmt.Println(op3)
	return nil
}

func Filter[V any](seq iter.Seq[V], ff func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for val := range seq {
			if ff(val) && !yield(val) { return }
		}
	}
}