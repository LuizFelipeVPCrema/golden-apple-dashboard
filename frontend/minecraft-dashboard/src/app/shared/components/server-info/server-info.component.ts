import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-server-info',
  templateUrl: './server-info.component.html',
  styleUrls: ['./server-info.component.scss'],
  imports: [CommonModule]
})
export class ServerInfoComponent {
  @Input() server: any;
}
