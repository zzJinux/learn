import "@typespec/http";
import "@typespec/rest";

using TypeSpec.Http;
using TypeSpec.Rest;

@service({
  title: "Managed PostgreSQL service",
})
namespace ManagedPostgreSQL;

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
}

model InstanceGroup {
  apiVersion: string;
  name: string;
  resourceVersion: int64;
  generation: int64;
  spec: InstanceGroupSpec;
  status: InstanceGroupStatus;
  phase: InstanceGroupPhase;
  nextInstanceOrdinal: int64;
}

model InstanceGroupSpec {
  compute: InstanceGroupComputeSpec;
}

model InstanceGroupComputeSpec {
    flavorId: string;
}

model InstanceGroupStatus {
  endpoints: string[];
}

enum InstanceGroupPhase {
    Provisioning,
    Active,
}