using Pulumi;
using Pulumi.Tls;

class MyStack : Stack
{
    public MyStack()
    {
        var key = new PrivateKey("my-private-key", new PrivateKeyArgs{
            Algorithm = "ECDSA",
            EcdsaCurve = "P384",
        });
    }
}
