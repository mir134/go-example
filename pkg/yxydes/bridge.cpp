#include "bridge.h"
#include "yxyDES.h"
using namespace std;
 yxyDES des;
void call() {
    string s="192.168.10.125";
	des.InitializeKey("zzm12345");
//	des.EncryptAnyLength(s.c_str());
	des.DecryptAnyLength("196eb7571ece01810b2b914708feb9fa");
	char* szDesIP = des.GetPlaintextAnyLength();
	printf(szDesIP);
}

char* Encrypt(char* szSource , char* key){
	des.InitializeKey(key);
	des.EncryptAnyLength(szSource);
	char* szDesIP = des.GetCiphertextAnyLength();
    return szDesIP;
}
char* Decrypt(char* szSource , char* key){
    des.InitializeKey(key);
    des.DecryptAnyLength(szSource);
    char* szDesIP = des.GetPlaintextAnyLength();
    return szDesIP;
}