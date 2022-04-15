"""A Python Pulumi program"""

import pulumi_tls as tls

key = tls.PrivateKey("my-private-key",
                     algorithm="ECDSA",
                     ecdsa_curve="P384"
                     )
