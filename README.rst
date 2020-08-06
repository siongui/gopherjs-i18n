===============
GopherJS_ i18n_
===============

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gopherjs-i18n?status.svg
   :target: https://godoc.org/github.com/siongui/gopherjs-i18n

.. image:: https://travis-ci.org/siongui/gopherjs-i18n.svg?branch=master
    :target: https://travis-ci.org/siongui/gopherjs-i18n

.. image:: https://gitlab.com/siongui/gopherjs-i18n/badges/master/pipeline.svg
    :target: https://gitlab.com/siongui/gopherjs-i18n/-/commits/master

.. image:: https://goreportcard.com/badge/github.com/siongui/gopherjs-i18n
   :target: https://goreportcard.com/report/github.com/siongui/gopherjs-i18n

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/gopherjs-i18n/blob/master/UNLICENSE


`gettext function`_ in your browser.

This package includes offline code to convert PO_ to JSON_,
and runtime code to translate strings.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.12.17`_
  - GopherJS_


Install
+++++++

Install package:

.. code-block:: bash

  go get -u github.com/siongui/gopherjs-i18n


Online Translating Website
++++++++++++++++++++++++++

See demo_ first. Demo code is in `example directory <example/>`_.

Wrap the string you want to translate in element with *data-default-string*
attribute containing the un-translated string. For example,

.. code-block:: html

  <div data-default-string="Home">Home</div>

or

.. code-block:: html

  <span data-default-string="Home">Home</span>

Both are valid for later translation.


Offline Preparing Translation Data
++++++++++++++++++++++++++++++++++

offline: example for converting PO_ to JSON_:

.. code-block:: go

  package main

  import "github.com/siongui/gopherjs-i18n/tool"

  func main() {
          po2json.PO2JSON("messages", "../pali/common/locale/", "po.json")
  }

*PO2JSON* takes three arguments:

  - *domain*: usually **messages**

  - *localedir*: the directory where you put PO_ files

  - *jsonPath*: output path of JSON_ file


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `[Golang] gettext Function on Frontend (Browser) by GopherJS <https://siongui.github.io/2016/01/28/go-gettext-function-frontend-browser-by-gopherjs/>`_

.. [2] `[Golang] GopherJS Synonyms with JavaScript <https://siongui.github.io/2016/01/29/go-gopherjs-synonyms-with-javascript/>`_

.. [3] `golang list file in directory <https://www.google.com/search?q=golang+list+file+in+directory>`_

       `ReadDir - ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/#ReadDir>`_

.. [4] `golang path join <https://www.google.com/search?q=golang+path+join>`_

       `path - The Go Programming Language <https://golang.org/pkg/path/>`_


.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _i18n: https://www.google.com/search?q=i18n
.. _gettext function: http://linux.die.net/man/3/gettext
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _JSON: http://www.json.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.12.17: https://golang.org/dl/
.. _demo: https://siongui.github.io/gopherjs-i18n/
.. _UNLICENSE: https://unlicense.org/
