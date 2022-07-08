// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.CertRequestArgs;
import com.pulumi.tls.Utilities;
import com.pulumi.tls.inputs.CertRequestState;
import com.pulumi.tls.outputs.CertRequestSubject;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new CertRequest(&#34;example&#34;, CertRequestArgs.builder()        
 *             .privateKeyPem(Files.readString(&#34;private_key.pem&#34;))
 *             .subject(CertRequestSubjectArgs.builder()
 *                 .commonName(&#34;example.com&#34;)
 *                 .organization(&#34;ACME Examples, Inc&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="tls:index/certRequest:CertRequest")
public class CertRequest extends com.pulumi.resources.CustomResource {
    /**
     * The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    @Export(name="certRequestPem", type=String.class, parameters={})
    private Output<String> certRequestPem;

    /**
     * @return The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    public Output<String> certRequestPem() {
        return this.certRequestPem;
    }
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="dnsNames", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> dnsNames() {
        return Codegen.optional(this.dnsNames);
    }
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="ipAddresses", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> ipAddresses() {
        return Codegen.optional(this.ipAddresses);
    }
    /**
     * Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
     * and ignored, as the key algorithm is now inferred from the key.
     * 
     * @deprecated
     * This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
     * 
     */
    @Deprecated /* This is now ignored, as the key algorithm is inferred from the `private_key_pem`. */
    @Export(name="keyAlgorithm", type=String.class, parameters={})
    private Output<String> keyAlgorithm;

    /**
     * @return Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
     * and ignored, as the key algorithm is now inferred from the key.
     * 
     */
    public Output<String> keyAlgorithm() {
        return this.keyAlgorithm;
    }
    /**
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     * 
     */
    @Export(name="privateKeyPem", type=String.class, parameters={})
    private Output<String> privateKeyPem;

    /**
     * @return Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
     * based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    @Export(name="subject", type=CertRequestSubject.class, parameters={})
    private Output</* @Nullable */ CertRequestSubject> subject;

    /**
     * @return The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
     * based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    public Output<Optional<CertRequestSubject>> subject() {
        return Codegen.optional(this.subject);
    }
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="uris", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> uris() {
        return Codegen.optional(this.uris);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertRequest(String name) {
        this(name, CertRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertRequest(String name, CertRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertRequest(String name, CertRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/certRequest:CertRequest", name, args == null ? CertRequestArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertRequest(String name, Output<String> id, @Nullable CertRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/certRequest:CertRequest", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CertRequest get(String name, Output<String> id, @Nullable CertRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertRequest(name, id, state, options);
    }
}
