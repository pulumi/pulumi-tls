// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SelfSignedCertSubject {
    /**
     * @return Distinguished name: `CN`
     * 
     */
    private @Nullable String commonName;
    /**
     * @return Distinguished name: `C`
     * 
     */
    private @Nullable String country;
    /**
     * @return Distinguished name: `L`
     * 
     */
    private @Nullable String locality;
    /**
     * @return Distinguished name: `O`
     * 
     */
    private @Nullable String organization;
    /**
     * @return Distinguished name: `OU`
     * 
     */
    private @Nullable String organizationalUnit;
    /**
     * @return Distinguished name: `PC`
     * 
     */
    private @Nullable String postalCode;
    /**
     * @return Distinguished name: `ST`
     * 
     */
    private @Nullable String province;
    /**
     * @return Distinguished name: `SERIALNUMBER`
     * 
     */
    private @Nullable String serialNumber;
    /**
     * @return Distinguished name: `STREET`
     * 
     */
    private @Nullable List<String> streetAddresses;

    private SelfSignedCertSubject() {}
    /**
     * @return Distinguished name: `CN`
     * 
     */
    public Optional<String> commonName() {
        return Optional.ofNullable(this.commonName);
    }
    /**
     * @return Distinguished name: `C`
     * 
     */
    public Optional<String> country() {
        return Optional.ofNullable(this.country);
    }
    /**
     * @return Distinguished name: `L`
     * 
     */
    public Optional<String> locality() {
        return Optional.ofNullable(this.locality);
    }
    /**
     * @return Distinguished name: `O`
     * 
     */
    public Optional<String> organization() {
        return Optional.ofNullable(this.organization);
    }
    /**
     * @return Distinguished name: `OU`
     * 
     */
    public Optional<String> organizationalUnit() {
        return Optional.ofNullable(this.organizationalUnit);
    }
    /**
     * @return Distinguished name: `PC`
     * 
     */
    public Optional<String> postalCode() {
        return Optional.ofNullable(this.postalCode);
    }
    /**
     * @return Distinguished name: `ST`
     * 
     */
    public Optional<String> province() {
        return Optional.ofNullable(this.province);
    }
    /**
     * @return Distinguished name: `SERIALNUMBER`
     * 
     */
    public Optional<String> serialNumber() {
        return Optional.ofNullable(this.serialNumber);
    }
    /**
     * @return Distinguished name: `STREET`
     * 
     */
    public List<String> streetAddresses() {
        return this.streetAddresses == null ? List.of() : this.streetAddresses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SelfSignedCertSubject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String commonName;
        private @Nullable String country;
        private @Nullable String locality;
        private @Nullable String organization;
        private @Nullable String organizationalUnit;
        private @Nullable String postalCode;
        private @Nullable String province;
        private @Nullable String serialNumber;
        private @Nullable List<String> streetAddresses;
        public Builder() {}
        public Builder(SelfSignedCertSubject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commonName = defaults.commonName;
    	      this.country = defaults.country;
    	      this.locality = defaults.locality;
    	      this.organization = defaults.organization;
    	      this.organizationalUnit = defaults.organizationalUnit;
    	      this.postalCode = defaults.postalCode;
    	      this.province = defaults.province;
    	      this.serialNumber = defaults.serialNumber;
    	      this.streetAddresses = defaults.streetAddresses;
        }

        @CustomType.Setter
        public Builder commonName(@Nullable String commonName) {
            this.commonName = commonName;
            return this;
        }
        @CustomType.Setter
        public Builder country(@Nullable String country) {
            this.country = country;
            return this;
        }
        @CustomType.Setter
        public Builder locality(@Nullable String locality) {
            this.locality = locality;
            return this;
        }
        @CustomType.Setter
        public Builder organization(@Nullable String organization) {
            this.organization = organization;
            return this;
        }
        @CustomType.Setter
        public Builder organizationalUnit(@Nullable String organizationalUnit) {
            this.organizationalUnit = organizationalUnit;
            return this;
        }
        @CustomType.Setter
        public Builder postalCode(@Nullable String postalCode) {
            this.postalCode = postalCode;
            return this;
        }
        @CustomType.Setter
        public Builder province(@Nullable String province) {
            this.province = province;
            return this;
        }
        @CustomType.Setter
        public Builder serialNumber(@Nullable String serialNumber) {
            this.serialNumber = serialNumber;
            return this;
        }
        @CustomType.Setter
        public Builder streetAddresses(@Nullable List<String> streetAddresses) {
            this.streetAddresses = streetAddresses;
            return this;
        }
        public Builder streetAddresses(String... streetAddresses) {
            return streetAddresses(List.of(streetAddresses));
        }
        public SelfSignedCertSubject build() {
            final var o = new SelfSignedCertSubject();
            o.commonName = commonName;
            o.country = country;
            o.locality = locality;
            o.organization = organization;
            o.organizationalUnit = organizationalUnit;
            o.postalCode = postalCode;
            o.province = province;
            o.serialNumber = serialNumber;
            o.streetAddresses = streetAddresses;
            return o;
        }
    }
}
