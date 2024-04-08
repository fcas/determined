.. _set-environment-images:

########################
 Set Environment Images
########################

Determined launches workloads using Docker containers. By default, workloads execute inside a
Determined-provided container that includes common deep learning libraries and frameworks. The
default containers can be found on the Determined Docker Hub with tags for each Determined version:

-  `Default containers with PyTorch <https://hub.docker.com/r/determinedai/pytorch-ngc>`__
-  `Containers for TensorFlow <https://hub.docker.com/r/determinedai/tensorflow-ngc>`__

By default, Determined will use the tag corresponding to your cluster's version. To specify a
different image from this default, update your job configuration to include:

.. code:: bash

   environment:
     image: # full CPU image path, e.g., determined/pytorch-ngc:<tag>

If one of the images above contain your required libraries, there is no additional environment
preparation needed.

If you need to add additional customization to the training environment, review the
:ref:`custom-env` page.
