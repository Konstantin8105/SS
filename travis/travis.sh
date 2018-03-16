#!/bin/bash

set -e

OUTFILE=/tmp/out.txt

function cleanup {
    EXIT_STATUS=$?

    if [ $EXIT_STATUS != 0 ]; then
        [ ! -f $OUTFILE ] || cat $OUTFILE
    fi

    exit $EXIT_STATUS
}
trap cleanup EXIT

if [ -f "$OUTFILE" ]; then
	rm -f $OUTFILE
fi

COVERAGE_FILES=`ls -1 *.coverprofile 2>/dev/null | wc -l`
if [ $COVERAGE_FILES != 0 ]; then
	rm -f *.coverprofile
fi

if [ -f coverage.txt ]; then
	rm -f coverage.txt
fi

export PKGS=$(go list ./... | grep -v /vendor/)

# Make comma-separated.
export PKGS_DELIM=$(echo "$PKGS" | paste -sd "," -)

# Run tests and append all output to out.txt. It's important we have "-v" so
# that all the test names are printed. It's also important that the covermode be
# set to "count" so that the coverage profiles can be merged correctly together
# with gocovmerge.
#
# Exit code 123 will be returned if any of the tests fail.
go list -f 'go test -v -tags integration -race -covermode atomic -coverprofile {{.Name}}.coverprofile -coverpkg $PKGS_DELIM {{.ImportPath}}' $PKGS | xargs -I{} bash -c "{} >> $OUTFILE"

# Merge coverage profiles.
COVERAGE_FILES=`ls -1 *.coverprofile 2>/dev/null | wc -l`
echo "Amount coverage files : ${COVERAGE_FILES}"
if [ $COVERAGE_FILES != 0 ]; then
	# check program `gocovmerge` is exist
	if which gocovmerge >/dev/null 2>&1; then
		gocovmerge `ls *.coverprofile` > coverage.txt
	fi
	rm -f *.coverprofile
fi

# Print stats
UNIT_TESTS=$(grep "=== RUN" $OUTFILE | wc -l | tr -d '[:space:]')
INT_TESTS=$(grep "# Total tests" $OUTFILE | cut -c21- | tr -d '[:space:]')

echo "Unit tests: ${UNIT_TESTS}"
echo "Integration tests: ${INT_TESTS}"

cat $OUTFILE

if [ -f coverage.txt ]; then
	rm -f coverage.txt
fi
