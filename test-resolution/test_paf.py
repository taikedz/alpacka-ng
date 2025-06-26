#!/usr/bin/env python3

import os
import shlex
import unittest
import subprocess


WHICH = "which which"

class Pman:
    def __init__(self, pman):
        self.envars = os.environ.copy()
        self.envars["PAF_TEST_PMAN"] = pman


    def run(self, cmd):
        res = subprocess.run(cmd, capture_output=True, env=self.envars)
        return res.returncode, res.stdout, res.stderr


    def ran_with_outputs(self, cmd, outputs=None, which=True):
        if which:
            outputs = [WHICH, WHICH] + outputs
        code, stdout, stderr = self.run(shlex.split(cmd))
        assert code == 0

        stdout_lines = [L for L in stdout.decode("utf-8").split("\n") if L]
        assert stdout_lines == outputs


class TestPaf(unittest.TestCase):
    def test_apt(self):
        pman = Pman("apt-get")
        pman.ran_with_outputs("bin/paf -u", ["sudo apt-get update"])
        pman.ran_with_outputs("bin/paf minetest", ["apt-cache search minetest"])
        pman.ran_with_outputs("bin/paf minetest -s", ["apt-cache show minetest"])
        pman.ran_with_outputs("bin/paf htop -sic vim", [
            "sudo apt-get install htop vim",
            "sudo apt-get autoclean",
            "sudo apt-get autoremove",
            ])
        pman.ran_with_outputs("bin/paf htop -is vim", ["apt-cache show htop", "apt-cache show vim"])
        pman.ran_with_outputs("bin/paf -iu tmux htop", ["sudo apt-get update", "sudo apt-get install tmux htop"])
        pman.ran_with_outputs("bin/paf -iug nginx", ["sudo apt-get update", "sudo apt-get upgrade"])
        pman.ran_with_outputs("bin/paf -ru tmux htop", ["sudo apt-get update", "sudo apt-get remove tmux htop"])

        pman.ran_with_outputs("bin/paf -x fix", ["sudo apt-get -f install"])
        pman.ran_with_outputs("bin/paf -x ppa=minetest/minetest",
            ["which add-apt-repository",
            "sudo add-apt-repository minetest/minetest"])
        pman.ran_with_outputs("bin/paf -x desc htop vim",
        [
            "--- htop ---", "apt-cache show htop",
            "--- vim ---", "apt-cache show vim",
        ])