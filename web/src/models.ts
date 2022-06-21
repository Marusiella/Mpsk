export interface Task {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    Name: string;
    Description: string;
    Hidden: boolean;
  }
  
  export interface Group {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    Name: string;
    Users?: any;
    Tasks: Task[];
  }
  export interface User {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    LastLoginTime: Date;
    Name: string;
    Email: string;
    Password: string;
    Role: number;
    Surname: string;
}