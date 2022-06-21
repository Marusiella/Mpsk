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