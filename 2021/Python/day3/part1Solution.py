

def main(lines):
    total_bits_per = len(lines[0].replace("\n", ""))
    bit_counts = [[0, 0] for bit in range(total_bits_per)]

    for line in lines:
        for index, bit in enumerate(line.replace('\n', '')):
            bit_val = int(bit)
            bit_counts[index][bit_val] += 1

    gamma_str = ""
    for count in bit_counts:
        if count[0] > count[1]:
            gamma_str += "0"
        else:
            gamma_str += "1"

    gamma_val = int(gamma_str, 2)

    mask = "1" * total_bits_per

    epsilon_val = gamma_val ^ int(mask, 2)

    print(bit_counts)
    print("gamma", gamma_val)
    print("epsilon", epsilon_val)
    print("power", gamma_val * epsilon_val)
