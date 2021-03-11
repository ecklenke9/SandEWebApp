import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';
import {Observable} from "rxjs";

@Injectable()
export class TodoService {
  constructor(private httpClient: HttpClient) {}

  getTodoList(): Observable<Todo[]>{
    return this.httpClient.get<Todo[]>(environment.gateway + '/todo');
  }

  addTodo(todo: Todo) {
    return this.httpClient.post(environment.gateway + '/todo', todo);
  }

  completeTodo(todo: Todo) {
    return this.httpClient.put(environment.gateway + '/todo', todo);
  }

  deleteTodo(todo: Todo) {
    return this.httpClient.delete(environment.gateway + '/todo/' + todo.id);
  }
}

export class Todo {
  id: string = "";
  message: string = "";
  complete: boolean = false;
}
