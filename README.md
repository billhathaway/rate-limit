rate-limit
--
limit the rate at which lines are copied from stdin to stdout

The typical use for this would be when piping between commands where part of the pipeline is expensive and should be throttled.

WARNING: end-of-line sequences are replaced with '\n', this might pose a problem if using lines ending with '\r\n'

Example:  

To remove all files from a huge directory, limit them to 100 per second to not cause all disk IOPs to be used.

    find /path/to/dir | rate-limit -l 100 | xargs rm

Usage:  

    rate-limit -l <lines_per_second>
