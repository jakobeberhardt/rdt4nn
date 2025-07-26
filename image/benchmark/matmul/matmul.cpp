#include <chrono>
#include <fstream>
#include <iostream>
#include <random>
#include <vector>

int main(int argc, char* argv[]) {
    if (argc != 2) { std::cerr << "Usage: " << argv[0] << " input.in\n"; return 1; }

    std::ifstream fin(argv[1]);
    size_t n, loops;
    if (!(fin >> n >> loops)) { std::cerr << "Bad input file\n"; return 1; }

    std::vector<double> A(n * n), B(n * n), C(n * n, 0.0);
    std::mt19937_64 rng(42);
    std::uniform_real_distribution<double> dist(0.0, 1.0);
    for (auto& v : A) v = dist(rng);
    for (auto& v : B) v = dist(rng);

    auto start = std::chrono::high_resolution_clock::now();
    for (size_t l = 0; l < loops; ++l) {
        for (size_t i = 0; i < n; ++i)
            for (size_t k = 0; k < n; ++k) {
                double a = A[i * n + k];
                for (size_t j = 0; j < n; ++j)
                    C[i * n + j] += a * B[k * n + j];
            }
    }
    auto end = std::chrono::high_resolution_clock::now();
    double seconds = std::chrono::duration<double>(end - start).count();
    double gflops = (2.0 * n * n * n * loops) / (seconds * 1e9);
    std::cout << "Elapsed: " << seconds << " s,  GFLOP/s: " << gflops << '\n';
    double checksum = 0.0;
    for (double v : C) checksum += v;
    std::cout << "Checksum: " << checksum << '\n';
    return 0;
}
