docker build -f Dockerfile-tests -t eval-multi-base-tests:latest .
docker run --rm eval-multi-base-tests:latest /home/runner/do/run_tests.sh

@rem docker run --rm -it eval-multi-base-tests:latest /bin/bash

