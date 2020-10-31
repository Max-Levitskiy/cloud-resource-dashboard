import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {ResourceService} from '../../../services/resource.service';

@Component({
  selector: 'app-resource-count',
  templateUrl: './resource-count.component.html',
  styleUrls: ['./resource-count.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ResourceCountComponent implements OnInit {

  constructor(private resourceService: ResourceService) { }

  resourceCount$;

  ngOnInit(): void {
    this.resourceCount$ = this.resourceService.fetchResourceCount();
  }
}
