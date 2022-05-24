cd src && pwd

min_coverage=90

ignored_files='Tests'

cov_profile=coverage.profile
cov_out=coverage.out
lcov_file=lcov.info

echo "Running coverage test with min_coverage=$min_coverage and ignored_files=$ignored_files"

cov_files=`go list ./... | grep -i -v $ignored_files`

cov_res="`go test -cover -coverprofile=$cov_profile $cov_files`"

echo "$cov_res"

cat $cov_profile | grep -i -v $ignored_files > $cov_out

gcov2lcov -infile=$cov_out -outfile=$lcov_file

if echo $cov_res | grep -i -q 'no test files'; then
    echo "ERR: All modules should be tested."
    exit 1
fi

if echo $cov_res | grep -i -q 'FAIL'; then
    echo "ERR: All modules should pass the tests."
    exit 1
fi

coverages=`echo $cov_res | grep -oP 'coverage\: (\d+)' | grep -oP '\d+'`

for cov in $coverages; do
    if ((cov < min_coverage)); then
        echo "ERR: All modules should have a coverage above $min_coverage."
        exit 1
    fi
done

