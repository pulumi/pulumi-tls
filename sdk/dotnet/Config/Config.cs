// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Tls
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("tls");

        private static readonly __Value<Pulumi.Tls.Config.Types.Proxy?> _proxy = new __Value<Pulumi.Tls.Config.Types.Proxy?>(() => __config.GetObject<Pulumi.Tls.Config.Types.Proxy>("proxy"));
        /// <summary>
        /// Proxy used by resources and data sources that connect to external endpoints.
        /// </summary>
        public static Pulumi.Tls.Config.Types.Proxy? Proxy
        {
            get => _proxy.Get();
            set => _proxy.Set(value);
        }

        public static class Types
        {

             public class Proxy
             {
                public bool? FromEnv { get; set; }
                public string? Password { get; set; } = null!;
                public string? Url { get; set; } = null!;
                public string? Username { get; set; } = null!;
            }
        }
    }
}
