# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = "centos/7"
  config.vm.box_check_update = false

  config.vm.provision "shell", inline: <<-SHELL
echo "Disabling SELINUX"
sudo tee /etc/selinux/config <<-'EOF'
SELINUX=disabled
SELINUXTYPE=targeted
EOF

echo "Setting hostname"
sudo tee /etc/sysconfig/network <<-'EOF'
NETWORKING=yes
HOSTNAME=gotutorial
EOF

echo "Getting Git"
yum install -y git

echo "Getting Go"
curl -o /tmp/go.tgz https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz

echo "Extracting Go"
tar -C /usr/local -zxf /tmp/go.tgz
rm /tmp/go.tgz

echo "Setting up Go environment"
sudo tee /etc/profile.d/go.sh <<-'EOF'
export GOROOT=/usr/local/go

if [[ ${PATH} != *"${GOROOT}/bin"* ]]; then
  export PATH=${PATH}:${GOROOT}/bin
fi
EOF

SHELL
  config.vm.provision :reload
  config.vm.provision "shell", inline: <<-SHELL
echo "Setting tutorial enviornment"
su - vagrant -c "git clone https://github.com/cmceniry/gotutorial.git tutorial"
tee /home/vagrant/.bashrc <<-'EOF'
# .bashrc

# Source global definitions
if [ -f /etc/bashrc ]; then
	. /etc/bashrc
fi

# User specific aliases and functions
export GOPATH=/home/vagrant/tutorial/workspace
EOF
chmod 0600 /home/vagrant/.bashrc
chown vagrant:vagrant /home/vagrant/.bashrc

echo "Done"
SHELL
end
