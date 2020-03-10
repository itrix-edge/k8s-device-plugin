__global__ void add(int *deviceCount);

extern "C" { int DeviceGetCount(void) {
    int deviceCount;
    cudaGetDeviceCount(&deviceCount);
    return deviceCount;
};}
