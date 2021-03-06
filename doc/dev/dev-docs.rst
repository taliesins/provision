.. Copyright (c) 2017 RackN Inc.
.. Licensed under the Apache License, Version 2.0 (the "License");
.. Digital Rebar Provision documentation under Digital Rebar master license
.. index::
  pair: Digital Rebar Provision; Documentation

.. _rs_dev_docs:

Developing Documentation
========================

As an open source project, we encourage community feedback and involvement.  Docs can be updated by
pull requests against the `github <https://github.com/digitalrebar/provision>`_ repository either from a private
tree or directly against the tree.

A couple of notes about consistency.

#. Digital Rebar is the name of the parent project and can be abbreviated DR.
#. Digital Rebar Provision or DR Provision or DRP can be used to reference this part of the project.
#. API docs generated from the go files as part of swagger annotations of the godoc comments.  Update there, please.
#. CLI docs generated from the cli files as part of cobra structures.  The tools generate those.  Update there, please.

Otherwise, try and find a good place for what needs to be added.  And Thanks!


Documentation Tooling
---------------------

There are a lot of ways to work with ReStructuredText...  Below is only one possible way of setting up a working environment to ensure you are writing clean RST based documentation.  This method is designed to be as lightlweight as possible, while still being as accurate as possible with the final rendered doc changes.   There are a lot of editors tha will render RST formatted markup, but very few of them render it the same way, or similar enough to the final rendered documentation styles to be correct.

This process uses the following elements:

#. Install Sphinx to correctly render the RST markup text to HTML
#. Edit the doc in any text editor of your preference (eg *vim*, *atom*, *emacs*, etc.)
#. Use a terminal window with a ``watch`` function to rebuild the HTML document tree
#. Use a web browser with an auto-refresh extension to view the rendered HTML


Install and Setup Sphinx and Swagger
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

These setup steps were tested and verified on a Mac platform.  However, they should be the same for Linux. 

1. Pre-requirement is to have ``pip`` installed:

  ::

      sudo easy_install pip


2. Install Sphinx:

  :: 

    pip install Sphinx

.. note:: for more detailed information, please review the Sphinx website on how to install it, at:
  http://www.sphinx-doc.org/en/stable/tutorial.html#install-sphinx


3. Install SwaggerDoc Library 

  ::

    sudo -H pip install -r requirements.txt 

.. note:: If you receive an error message on your first HTML tree build (when you run ``make html``) similar to:
    ::

       sphinx-build -b html -d _build/doctrees   . _build/html
       Running Sphinx v1.6.5
       Extension error:
       Could not import extension sphinxcontrib.swaggerdoc (exception: No module named swaggerdoc)
       make: *** [html] Error 1


  You will need to build and install the *swaggerdoc* components manually.  To do so, do the following:
    ::

      git clone https://github.com/galthaus/sphinx-swaggerdoc
      cd sphinx-swaggerdoc/
      sudo python setup.py  install
      cd ..
      rm -rf spinx-swaggerdoc/


Edit the Docs
~~~~~~~~~~~~~

  * Checkout/clone the Digital Rebar Provision repo from Github
  * Modify the doc(s) as appropriate
  * Verify the modifications are rendered correctly 
  * Create a branch
  * Submit a pull request for your changes


Rebuild the HTML Rendered Docs
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

To assist with editing/reviewing your changes, you will want to rebuild the RST format documentation to rendered HTMl.  The Sphinx package is used to do this.  Generally, you simply change to the *base git cloned Provision* directory, and do:

  ::

    make html

If you are making a lot of changes, you will want to use a file watching utility to see watch for file writes, and automatically kick the ``make html`` process off for you.

On Mac OS X - you can install the ``fswatch`` package:

  ::

    brew install fswatch

An example use would be:

  ::

      # cd to the base Provision git directory
      fswatch -o -0 -r doc | xargs -0 -n 1 -I {} make html

Now, any time any changes are made to the files in the ``doc/`` directory, the ``make html`` process will be automatically kicked off for you.  You should run this in a separate terminal window.


View Your Doc Edits Locally
~~~~~~~~~~~~~~~~~~~~~~~~~~~

Once you've run ``make html`` as above, the _RSTr_ format files will be built into _HTML_ generated versions.  You can view the effects of your changes by opening your web browser up, and navigating to the _HTML_ generated docs on your local disk.  The below path example references my home directory, and github path location.  You'll have to modify this to your local User and location where you've stored your github clone.

Point your browser to the on-disk rendered location (which is the ``_build/`` directory in the base git repo on disk).  For example:

  ``file:///Users/shane/github/digitalrebar/provision/_build/html/doc/dev/dev-docs.html``


Auto-Refresh Browser
~~~~~~~~~~~~~~~~~~~~

The last piece of the puzzle, you will want to set your web browser to auto-refresh a given tab or window.  This way, the HTML rendered documentation will be refreshed in the browser.   There are several add-ons/extensions that will do this for you.  Here at RackN we have used the following extensions:

  Chrome *Auto Refresh Plus* extension:
    https://chrome.google.com/webstore/detail/auto-refresh-plus/hgeljhfekpckiiplhkigfehkdpldcggm

  Firefox *Tab Reloader* add-on (works on Chrome, Firefox, and Opera; but limited to 10 second reloads as minimum reload time):
    https://add0n.com/tab-reloader.html

Simply set your browser tab to refresh every 5 or so seconds.

