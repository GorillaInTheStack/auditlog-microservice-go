# Vagrant & Ansible
This is a quick way to deploy the microservice into a VM (currently set to be vbox VM) and configure it using ansible.

Run
```bash
vagrant up
```
To bring up the VM.
The ansible playbook will run automatically once the VM is running.
The playbook will make sure to install git and golang in the VM.

You can run 
```bash 
ansible-playbook playbook.yml
```
To run the playbook on host.