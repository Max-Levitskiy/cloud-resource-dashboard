export class Resource {
    CloudProvider: string;
    ResourceType: string;
    ProjectId: string;
    Name: string;
    Region: string;
    CreationDate: Date;
    Tags: string[];
}

export class ResourceCount {
  count: number;
}
