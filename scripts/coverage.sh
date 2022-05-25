echo "Starting coverage test"

cd src && pwd

set -e

cov_profile=coverage.profile
cov_out=coverage.out
lcov_file=lcov.info

echo "Running courtney on all modules"

courtney -e ./...

echo "Creating '$lcov_file' from '$cov_out'"

gcov2lcov -infile=$cov_out -outfile=$lcov_file

echo "Completed coverage test "