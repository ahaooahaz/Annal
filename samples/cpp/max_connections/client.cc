#include <common.h>


int main() {
    int count = 0;

    sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr(ADDR);
    addr.sin_port = htons(PORT);

    while (true) {
        int fd = socket(AF_INET, SOCK_STREAM, 0);
        if (fd <= 0) {
            ERROR("create socket failed, ret code: %d\n", fd);
            return -1;
        }

        int ret = connect(fd, (sockaddr*)&addr, sizeof(addr));
        if (ret < 0) {
            ERROR("connect failed, ret code: %d\n", ret);
            return -1;
        }
        cout << count++ << ":connected" << endl;
    }

    return 0;
}