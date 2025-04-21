import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { ServerService } from '../../core/services/server.service';

@Component({
  selector: 'app-server-creator',
  templateUrl: './server-creator.component.html',
  styleUrls: ['./server-creator.component.scss'],
  imports: [CommonModule, FormsModule]
})
export class ServerCreatorComponent {
  server = {
    name: '',
    port: '25565',
    gameMode: 'survival',
    levelType: 'easy'
  };

  gameModes = ['survival', 'creative', 'adventure', 'spectator'];
  levelTypes = ['easy', 'medium', 'hard'];

  constructor(private router: Router, private serverService: ServerService) {}

  submitServer() {
    this.serverService.create(this.server).subscribe({
      next: () => {
        console.log('Servidor criado com sucesso!');
        this.router.navigate(['/dashboard']);
      },
      error: (err) => {
        console.error('Erro ao criar servidor:', err);
        alert('Falha ao criar servidor. Verifique os dados e tente novamente.');
      }
    });
  }
}