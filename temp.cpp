#include <iostream>
#include <vector>

class Solution {
public:
    int knightDialer(int n) {
        int ans = 0;
        std::vector<std::vector<int>> v;
        for (int i = 0; i < 3; i++) {
            std::vector<int> temp;
            for (int j = 1; j <= 3; j++) {
                temp.push_back(3 * i + j);
            }
            v.push_back(temp);
        }
        v.push_back({11, 0, 12});

        // std::cout << "Vector values:" << std::endl;
        // for (const auto &row : v) {
        //     for (int val : row) {
        //         std::cout << val << " ";
        //     }
        //     std::cout << std::endl;
        // }
                        ans %= mm;
                    }
                }
            }
            return ans;
        }

        int knightDialer(int n) {
            int ans = 0;
            for (int i = 0; i < 3; i++) {
                vector<int> temp;
                for (int j = 1; j <= 3; j++) {
                    temp.push_back(3 * i + j);
                }
                v.push_back(temp);
            }
            v.push_back({11, 0, 12});

            for (int i = 0; i < 4; i++) {
                for (int j = 0; j < 3; j++) {
                    ans += rec(i, j, n);
                    ans %= mm;
                }
            }
            return ans;
        }
    };

    int knightDialer(int n) {
        int ans = 0;
        for (int i = 0; i < 3; i++) {
            vector<int> temp;
            for (int j = 1; j <= 3; j++) {
                temp.push_back(3 * i + j);
            }
            v.push_back(temp);
        }
        v.push_back({11, 0, 12});

        // cout << "Vector values:" << endl;
        // for (const auto &row : v) {
        //     for (int val : row) {
        //         cout << val << " ";
        //     }
        //     cout << endl;
        // }


        for (int i = 0; i < 4; i++) {
            for (int j = 0; j < 3; j++) {  // Adjusted loop limit for j
                // ans += rec(i, j, n);
                // ans %= mm;
            }
        }
        rec(0,0,1);
        return ans;
    }
};
