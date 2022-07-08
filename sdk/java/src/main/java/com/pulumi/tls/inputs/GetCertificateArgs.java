// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCertificateArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCertificateArgs Empty = new GetCertificateArgs();

    @Import(name="content")
    private @Nullable Output<String> content;

    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    @Import(name="url")
    private @Nullable Output<String> url;

    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="verifyChain")
    private @Nullable Output<Boolean> verifyChain;

    public Optional<Output<Boolean>> verifyChain() {
        return Optional.ofNullable(this.verifyChain);
    }

    private GetCertificateArgs() {}

    private GetCertificateArgs(GetCertificateArgs $) {
        this.content = $.content;
        this.url = $.url;
        this.verifyChain = $.verifyChain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCertificateArgs $;

        public Builder() {
            $ = new GetCertificateArgs();
        }

        public Builder(GetCertificateArgs defaults) {
            $ = new GetCertificateArgs(Objects.requireNonNull(defaults));
        }

        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        public Builder content(String content) {
            return content(Output.of(content));
        }

        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public Builder verifyChain(@Nullable Output<Boolean> verifyChain) {
            $.verifyChain = verifyChain;
            return this;
        }

        public Builder verifyChain(Boolean verifyChain) {
            return verifyChain(Output.of(verifyChain));
        }

        public GetCertificateArgs build() {
            return $;
        }
    }

}