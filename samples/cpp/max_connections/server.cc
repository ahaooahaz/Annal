#include <sys/types.h>
#include <sys/socket.h>
#include <common.h>
#include <unistd.h>
using namespace std;

int main(int argc, char* argv[]) {
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if (fd <= 0) {
        ERROR("create socket failed, ret code: %d\n", fd);
        return -1;
    }

    sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr(ADDR);
    addr.sin_port = htons(PORT);

    int ret = bind(fd, (sockaddr*)&addr, sizeof(addr));
    if (ret == -1) {
        ERROR("bind addr failed, ret code: %d\n", ret);
        return -1;
    }

    if (listen(fd, 5) == -1) {
        ERROR("listen failed, err: %d\n", -1);
        return -1;
    }

    sockaddr peer_addr;
    socklen_t len;
    int count = 0;
    while (true) {
        int cfd = accept(fd, &peer_addr, &len);
        if (cfd == -1) {
            ERROR("appect failed, ret code: %d\n", -1);
            return -1;
        }

        cout << count++ << ":fd:" << cfd << endl;
    }
    return 0;
}