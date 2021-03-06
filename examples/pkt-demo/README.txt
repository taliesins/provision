# Copyright (c) 2018 RackN Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


README file for pkt-demo

NOTE:  This example used to be called "5min-demo".  It has been renamed
       to "pkt-demo" to more closely align with the use case, and 
       expanded scope.


OVERVIEW AND IMPORTANT REQUIREMENTS
-----------------------------------

You will need to perform the following FOUR EASY STEPS in preparation 
to using this pkt-demo process.  Details for each step are below.

  1. GIT CLONE
    - get the pkt-demo code from the github repo

  2. SECRETS
     - get your API and Username secrets - modify 'private-contents/secrets' file
       * have a packet.net account, and your API KEY, and PROJECT ID
       * get your RackN USERNAME for registered content download authorization

  3. CUSTOMIZE
     - make changes to the terraform "vars.tf" parameters

  4. RUN demo-run.sh SCRIPT
     - actually run the demo setup script

Additional Sections (hopefully useful) of documentation:
  
  * WHAT HAPPENS
  * PACKET SECRETS NOTE
  * GENERAL NOTES
  * CLEANUP
  * OPERATING AND TROUBLESHOOTING
  * ADVANCED USAGE OPTIONS

DETAILED STEPS
--------------

1. GIT CLONE
============

*If you do not have the Repo Cloned already, do:*

  The following steps will clone this content from the digitalrebar/provision 
  github repo (we assume you will run this from $HOME/pkt-demo - adjust yourself
  accordingly if you want to put it somewhere else):

    git clone -n https://github.com/digitalrebar/provision.git --depth=1
    cd provision
    git checkout HEAD examples/pkt-demo
    cp -r examples/pkt-demo $HOME/pkt-demo
    cd $HOME/pkt-demo

  DO NOT run these steps with 'sudo', your username should own the new directory and files.


*If you ALREADY have the repo cloned locally, do:*

  Simply copy the `digitalrebar/provision/examples/pkt-demo` directory to a new
  location.  For example: 

    cd <path_to_github_clone>/digitalrebar/provision/examples/
    git pull
    cp -r pkt-demo $HOME/pkt-demo
    cd $HOME/pkt-demo


2. SECRETS
==========

ABSOLUTELY NECESSARY - to authenticate to the packet.net API services and spin
up Nodes, and to authenticate to the RackN Portal to download content and
plugins.

EDIT THE SECRETS FILE !!  Located in 'private-content/secrets'.  You need:

  API       packet.net key for access to your Packet Project
  PROJECT   packet.net project to create DRP and Nodes in 
  USERNAME  your RackN username ID (UUID string)

    # modify the API, PROJECT, and USERNAME  variables

    vim private-content/secrets

  * API and PROJECT are from packet.net and you should find them in your
    Packet portal management (see PACKET SECRETS NOTE below for more details).

  * Contact RackN for demo credit if you don't have a packet account yet.
    - contact: http://www.rackn.com/support/slack/

  * USERNAME is from the RackN Portal - to find your USERNAME, log in 
    to the portal, and navigate to:

      Hamburger Menu (3 horizontal lines in upper left)
      User Profile
      Account ID

      Direct URL:  https://portal.rackn.io/#/user/

    It will be a big ugly UUID like string like:  ad9914b7-60bd-49d9-81d0-95e532e7ce1c


NOTE: Please do not modify the following in the 'secrets' file:
      API_KEY, PROJECT_ID, and RACKN_USERNAME 


3. CUSTOMIZATION
================

Make sure you've modified the 'secrets' file appropriately (see '3. SECRETS' above).

These steps are OPTIONAL - if you are just trying to "quickstart" and play - leave
these as is.  If you want more Machines and some customization, read on. 

These define what Packet resources you use, and thus; how much money you may end up 
owing Packet.   By default we ONLY spin up 1 type_0 machine for the Digital Rebar
Provision endpoint, and one type_0 test CentOS 7 machine provisioned against the DRP 
endpoint in the Packet EWR1 (Parsipany, NJ, USA) datacenter.  

Packet pricing:  https://www.packet.net/bare-metal/ 

  * endpoint_type
    - The DRP Endpoint node type in Packet parlance.  Type_0 is WAY more than you need.
    - see: https://www.packet.net/developers/api/plans/
    
  * drp_os
    - you can modify which Operating System the DRP endpoint is running on 
    - only supported ones are CentOS 7 and Ubuntu 16.04

  * machines_type
    - The DRP Endpoint Machines type; Packet calls these "plans".
    - see: https://www.packet.net/developers/api/plans/

  * machines_count
    - specify the number of Machines to provision (default is 1 machine of type_0)
    - these will cost you money/credits from Packet

  * machines_os
    - you can modify which Operating System the Machines will be provisioned with
    - only supported ones are CentOS 7 and Ubuntu 16.04

  * packet_facility
    - change the packet.net facility to provsion the cluster in  (default is EWR1)
    - see: https://www.packet.net/developers/api/facilities/


4. RUN demo-run.sh SCRIPT
=========================

The 'demo-run.sh' is the control script that will walk you through the deployment
process.  Simply start this script. 

    ./demo-run.sh

If you re-run the script and want to skip steps that have run previously, simply 
answer "no" to the "ACTION" input.  HIGHLY recommended to run this the first time
in interactive mode (just do the above script run).

Basic USAGE options for "demo-run.sh"

  CONFIRM=no ./demo-run.sh      # disable prompting for each step - auto run
                                # we suggest NOT doing this the first time!!

  SKIP_LOCAL=yes ./demo-run.sh  # skip installing DRP locally - if you have a
                                # current copy installed already - mostly used
                                # in bandwidth constrained environments to 
                                # avoid downloading the dr-provision.zip

  CONFIRM and SKIP_LOCAL can be combined if you choose


ADDITIONAL SECTIONS
-------------------


WHAT HAPPENS?
=============

1.  set PATH to include the ./bin directory for DRP and terraform/etc.
2.  install terraform locally in your ./bin directory 
    (can be skipped, eg for demo in bandwidth constrained environments)
3.  install the secrets to the terraform vars.tf file
4.  create SSH keys for DRP endpoint and Machines
5.  inject the SSH keys in to packet.net Project
6.  build the DRP Endpoint server in packet.net 
    (on centos-7, configurable, see vars.tf)
7.  install DRP locally in ./bin for 'drpcli' control
8.  install DRP on the remote endpoint 
9.  configure content on DRP endpoint [see: GENERAL NOTES: CONTENT]
10. set the DRP endpoint IP address for terraform 'drp-machines.tf'
    this modifies Packet node references to PXE off of DRP endpoint
11. boot "N" number of machines to provision against the new
    DRP endpoint (default is 1 Machine of type_0 plan)
    (set "N" in vars.tf for "machines_count" variable)

  
PACKET SECRETS NOTE
===================
Unfortunately, Packet doesn't provide a clean way we can point you to 
for finding your API and PROJECT identities.  You'll have to find these
from the portal, hopefully, following the below path (YMMV):

   * log in to:  https://app.packet.net/login 
   * use down arrow next to User icon in upper right, select "Api Keys"
   * select the API key you wish to use 
   * use the back arrow in the upper left of the Portal 
     (not your browser back button)
   * use the down arrow next to User icon, select "Change Organization"
   * select the Org/Project you want to place resources in 
   * use the down arrow next to User icon, select "Settings" underneath 
     Project name
   * that page has the Project ID (labeled "Organization ID")

CLEANUP
=======

You can cleanup/reset the pkt-demo/ directory back to "factory
defaults" with the following:

  bin/control.sh cleanup        # runs `terraform destroy`
                                # destroys SSH keys
                                # cleansup local directory artifacts

IF YOU WANT to rerun this process, we suggest you do 
  cp private-contents/secrets secrets
Then, to restore the secrets on the next run, do:
  cp secrets private-contents/

(before running 'demo-run.sh' ... You won't have to modify the 'secrets'
file in the future.

WARNING:  CLEANUP NUKES everything !!  EVERYTHING !! ...and resets the
          demo directory to the beginning


GENERAL NOTES
=============

 CONTENT  By default demo-run.sh will have the DRP endpoint
          download content from the endpoint - in many 
          situations - you'd download content to your laptop
          and push the content to the DRP endpoint (eg the
          endpoint has no direct internet connection). 

          To force "proxy" pushing content - call the run-demo.sh
          script with the "local" argument as ARGv2.  Note that
          this options _should_ work but is not very well 
          tested (hint: there are probably some minor bugs)


 PREREQS  The machine you run the 'demo-run.sh' from must meet a
          few basic requirements.  We support Mac OS X and Linux.
          The demo script will give you a message if you're missing
          anything important.  Here are a couple of notes on things
          you'll probably need if you don't have them already... 


          jq         tool for parsing JSON text, REQUIRED
                     Mac:   1. Install Homebrew (https://brew.sh)
                            2. 'brew install jq'
                     Linux: 'sudo yum -y install jq'
                        or: 'sudo apt -y install jq'

OPERATING AND TROUBLESHOOTING
=============================

Some quick pointers for using or troubleshooting the environment that
gets set up...   All of these steps assume you are in the base directory
that you ran the demo-run.sh script from.

  export CLUSTER_NAME=<the_name_you_set_for_your_cluster>

  export DRP_ID=`bin/control.sh get-drp-id`
  export DRP_ADDR=`bin/control.sh get-address $DRP_ID`

SSH to the DRP endpoint (uses above variables):

  ssh -x -i ./${CLUSTER_NAME}-drp-ssh-key root@$DRP_ADDR

SSH to a Machine (you must get the Machine IP address from DRP Endpoint):

  ssh -x -i ./${CLUSTER_NAME}-machines-ssh-key root@<_MACHINE_IP_ADDRESS>


Destroy everything - see the CLEANUP section above.  This destroys
the Packet nodes, so you no longer are charged for them.


ADVANCED USAGE OPTIONS
======================

"demo-run.sh" just drives the "bin/control.sh" script to make it easy 
and prettier.  You can run the full demo without any Confirmation prompts, 
set CONFIRM variable to "no":

  CONFIRM=no ./demo-run.sh

The entire demo will run through without (hopefully...) any interactions. 


If you are running in a bandwidth constrained environment (ge poor WiFi, 
or Cellular based Hotspot), and IF you have current version of DRP 
installed on your laptop (and in your PATH), then you can skip the local
download/install step, with:

  SKIP_LOCAL=yes ./demo-run.sh


CONFIRM and SKIP_LOCAL can be combined if you want.


You can manually drive some things with the "bin/control.sh" script - simply
run it with the "--usage" or "--help" flags, it'll print out usage statement. 


"bin/control.sh cleanup" has an 8 second safety timer in it.  If you know 
what you're doing - you can simply call it with "bin/control.sh cleanup 
force" - and it'll skip the safety timer. 


You can get your DRP Endpoint provisioned IP address with the "bin/control.sh" 
script (AFTER it has been successfully provisioned, of course):

  DRP_ID=`bin/control.sh get-drp-id`
  DRP_ADDR=`bin/control.sh get-address $DRP_ID`


You can SSH directly to the DRP Endpoint using the injected SSH keys:

  ssh -x -i ./drp-ssh-key root@$DRP_IP 
  OR
  bin/control.sh ssh $DRP_ID


You should be able to SSH to the Machines directly as well, using the following
(after they've been provsioned and installed, of course):

  terraform plan    # to get the various machines IP addresses
                    # or get from packet.net

  ssh -x -i ./machines-ssh-key root@<MACHINE_IP>
                  

IF you need a custom version of the terraform-provider-packet plugin, you can 
specify it in your ~/.terraformrc file like so:

  providers { packet = "$HOME/.terraform.d/terraform-provider-packet" }

Prior to running "terraform init" (which occurs during the "terraform-install"
stage in this Demo.
