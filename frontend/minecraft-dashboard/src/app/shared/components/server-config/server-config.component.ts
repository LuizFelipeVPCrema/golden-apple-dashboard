import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-server-config',
  templateUrl: './server-config.component.html',
  styleUrls: ['./server-config.component.scss']
})
export class ServerConfigComponent {
  @Input() server: any;

  // Aqui você pode emitir eventos para salvar ou atualizar configurações
}
