This is a eval-multi-base docker image.

It includes all the languages but not the runner.

We want to split those so that rebuilding the image with the runner
is faster.
