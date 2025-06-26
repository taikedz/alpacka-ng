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
        self.priors = ["which which"]


    def run(self, cmd):
        res = subprocess.run(cmd, capture_output=True, env=self.envars)
        return res.returncode, res.stdout, res.stderr


    def ran_with_outputs(self, cmd, outputs=None):
        outputs = self.priors + outputs
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

    def test_dnf(self):
        pman = Pman("dnf")
        pman.ran_with_outputs("bin/paf -u", [])
        pman.ran_with_outputs("bin/paf minetest", ["dnf search minetest"])
        pman.ran_with_outputs("bin/paf minetest -s", ["dnf info minetest"])
        pman.ran_with_outputs("bin/paf htop -sic vim", [
            "sudo dnf install htop vim",
            "sudo dnf clean",
            "sudo dnf autoremove",
            ])
        pman.ran_with_outputs("bin/paf htop -is vim", ["dnf info htop", "dnf info vim"])
        pman.ran_with_outputs("bin/paf -iu tmux htop", ["sudo dnf install tmux htop"])
        pman.ran_with_outputs("bin/paf -iug nginx", ["sudo dnf upgrade"])
        pman.ran_with_outputs("bin/paf -ru tmux htop", ["sudo dnf remove tmux htop"])

    def test_pacman(self):
        pman = Pman("pacman")
        pman.ran_with_outputs("bin/paf -u", ["sudo pacman -Sy"])
        pman.ran_with_outputs("bin/paf minetest", ["pacman -Ss minetest"])
        pman.ran_with_outputs("bin/paf minetest -s", ["pacman -Si minetest"])
        pman.ran_with_outputs("bin/paf htop -sicy vim", [
            "sudo pacman -S --noconfirm htop vim",
            "sudo pacman -Scc",
            ])
        pman.ran_with_outputs("bin/paf htop -is vim", ["pacman -Si htop", "pacman -Si vim"])
        pman.ran_with_outputs("bin/paf -iu tmux htop", ["sudo pacman -Sy", "sudo pacman -S tmux htop"])
        pman.ran_with_outputs("bin/paf -iug nginx", ["sudo pacman -Sy", "sudo pacman -Su"])
        pman.ran_with_outputs("bin/paf -ru tmux htop", ["sudo pacman -Sy", "sudo pacman -Rs tmux htop"])

    def test_zypper(self):
        pman = Pman("zypper")
        pman.ran_with_outputs("bin/paf -u", ["sudo zypper refresh"])
        pman.ran_with_outputs("bin/paf minetest", ["zypper search minetest"])
        pman.ran_with_outputs("bin/paf minetest -s", ["zypper info minetest"])
        pman.ran_with_outputs("bin/paf htop -sicy vim", [
            "sudo zypper install -y htop vim",
            "sudo zypper clean",
            ])
        pman.ran_with_outputs("bin/paf htop -is vim", ["zypper info htop", "zypper info vim"])
        pman.ran_with_outputs("bin/paf -iu tmux htop", ["sudo zypper refresh", "sudo zypper install tmux htop"])
        pman.ran_with_outputs("bin/paf -iug nginx", ["sudo zypper refresh", "sudo zypper update"])
        pman.ran_with_outputs("bin/paf -ru tmux htop", ["sudo zypper refresh", "sudo zypper remove tmux htop"])
