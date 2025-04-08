// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.tls.Utilities;
import com.pulumi.tls.inputs.GetCertificateArgs;
import com.pulumi.tls.inputs.GetCertificatePlainArgs;
import com.pulumi.tls.inputs.GetPublicKeyArgs;
import com.pulumi.tls.inputs.GetPublicKeyPlainArgs;
import com.pulumi.tls.outputs.GetCertificateResult;
import com.pulumi.tls.outputs.GetPublicKeyResult;
import java.util.concurrent.CompletableFuture;

public final class TlsFunctions {
    public static Output<GetCertificateResult> getCertificate() {
        return getCertificate(GetCertificateArgs.Empty, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetCertificateResult> getCertificatePlain() {
        return getCertificatePlain(GetCertificatePlainArgs.Empty, InvokeOptions.Empty);
    }
    public static Output<GetCertificateResult> getCertificate(GetCertificateArgs args) {
        return getCertificate(args, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetCertificateResult> getCertificatePlain(GetCertificatePlainArgs args) {
        return getCertificatePlain(args, InvokeOptions.Empty);
    }
    public static Output<GetCertificateResult> getCertificate(GetCertificateArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("tls:index/getCertificate:getCertificate", TypeShape.of(GetCertificateResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetCertificateResult> getCertificate(GetCertificateArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("tls:index/getCertificate:getCertificate", TypeShape.of(GetCertificateResult.class), args, Utilities.withVersion(options));
    }
    public static CompletableFuture<GetCertificateResult> getCertificatePlain(GetCertificatePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("tls:index/getCertificate:getCertificate", TypeShape.of(GetCertificateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPublicKeyResult> getPublicKey() {
        return getPublicKey(GetPublicKeyArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPublicKeyResult> getPublicKeyPlain() {
        return getPublicKeyPlain(GetPublicKeyPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPublicKeyResult> getPublicKey(GetPublicKeyArgs args) {
        return getPublicKey(args, InvokeOptions.Empty);
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPublicKeyResult> getPublicKeyPlain(GetPublicKeyPlainArgs args) {
        return getPublicKeyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPublicKeyResult> getPublicKey(GetPublicKeyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("tls:index/getPublicKey:getPublicKey", TypeShape.of(GetPublicKeyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPublicKeyResult> getPublicKey(GetPublicKeyArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("tls:index/getPublicKey:getPublicKey", TypeShape.of(GetPublicKeyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get a public key from a PEM-encoded private key.
     * 
     * Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.tls.PrivateKey;
     * import com.pulumi.tls.PrivateKeyArgs;
     * import com.pulumi.tls.TlsFunctions;
     * import com.pulumi.tls.inputs.GetPublicKeyArgs;
     * import com.pulumi.std.StdFunctions;
     * import com.pulumi.std.inputs.FileArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var ed25519_example = new PrivateKey("ed25519-example", PrivateKeyArgs.builder()
     *             .algorithm("ED25519")
     *             .build());
     * 
     *         // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
     *         final var privateKeyPem-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyPem(ed25519_example.privateKeyPem())
     *             .build());
     * 
     *         // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
     *         final var privateKeyOpenssh-example = TlsFunctions.getPublicKey(GetPublicKeyArgs.builder()
     *             .privateKeyOpenssh(StdFunctions.file(FileArgs.builder()
     *                 .input("~/.ssh/id_rsa_rfc4716")
     *                 .build()).result())
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPublicKeyResult> getPublicKeyPlain(GetPublicKeyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("tls:index/getPublicKey:getPublicKey", TypeShape.of(GetPublicKeyResult.class), args, Utilities.withVersion(options));
    }
}
