import "@typespec/http";
import "@typespec/rest";

using TypeSpec.Http;
using TypeSpec.Rest;

@service({
  title: "Managed PostgreSQL Service",
})
namespace Mapo;

@route("/instance-groups")
namespace InstanceGroups {
  op list(): {
    @statusCode statusCode: 200;
    @body instanceGroups: InstanceGroup[];
  };

  op create(@body instanceGroup: InstanceGroup): {
    @statusCode statusCode: 201;
    @body instanceGroup: InstanceGroup;
  };

  op read(@path name: string): {
    @statusCode statusCode: 200;
    @body instanceGroup: InstanceGroup;
  } | {
    @statusCode statusCode: 404;
  };

  @put
  op update(@path name: string, @body instanceGroup: InstanceGroup): {
    @statusCode statusCode: 200;
    @body instanceGroup: InstanceGroup;
  } | {
    @statusCode statusCode: 404;
  };

  @delete
  op delete(@path name: string): {
    @statusCode statusCode: 204;
  };
}

model TypeMeta {
  apiVersion: string;
}

model ObjectMeta {
  name: string;
  resourceVersion: int64;
  generation: int64;
  creationTimestamp?: utcDateTime;
  deletionTimestamp?: utcDateTime;
  labels: Record<string>;
  annotations: Record<string>;
  finalizers: string[];
}

model Scope {
  projectId: string;
  projectName: string;
  projectDomainId: string;
  projectDomainName: string;
}

model Creator {
  userId: string;
  userName: string;
  userDomainId: string;
  userDomainName: string;
  // sourceIpAddress?: string; // Uncomment if use-case required
  // userAgent?: string; // Uncomment if use-case required
  // role?: string; // Uncomment if role information is required
}

model SecurityGroup {
  id: string;
}

model Subnet {
  id: string;
  azName: string;
  // replicas?: int64; // Uncomment if replicas use-case required
}

model InstanceGroupNetworkSpec {
  securityGroups: SecurityGroup[];
  primarySubnet: Subnet;
  secondarySubnet: Subnet[];
}

model InstanceGroupSpec {
  // multiAzFlag?: boolean; // Uncomment if Multi-Az flag use-case required
  network: InstanceGroupNetworkSpec;
}

model InstanceGroupStatus {
  haState: string;
  state: string;
  endpoints: string[];
  isInitializing: boolean;
  isConfigured: boolean;
  dnsEnabled: boolean;
}

model InstanceGroup {
  typeMeta: TypeMeta;
  metadata: ObjectMeta;
  scope: Scope;
  creator: Creator;
  spec: InstanceGroupSpec;
  status: InstanceGroupStatus;
  nextInstanceOrdinal: int64;
}
