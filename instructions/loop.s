# prerequisite for instruction
li {r1}, {v1}
li {r2}, {v2}

# set pin high
sw s1, GPIO_OUTPUT_VAL(s0)

# run instruction
{instruction} {r3}, {r1}, {r2}

# set pin low
sw s1, GPIO_OUTPUT_VAL(s0)

