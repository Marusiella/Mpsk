import React, { useEffect, useState } from "react";
import { address } from "./main";
import ManageTasks from "./ManageTasks";
import { Group, User } from "./models";
import {
  AppShell,
  Navbar,
  Header,
  Footer,
  Aside,
  Text,
  MediaQuery,
  Burger,
  useMantineTheme,
  Center,
  Box,
  MantineTheme,
} from "@mantine/core";
export default function HomePage() {
  const [root, setRoot] = useState<Group[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  function Fetch() {
    fetch(address + "/api/v1/gettasks", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: Group[]) => {
        console.log([data]);
        try {
          data.map(() => {});
          setRoot(data);
        } catch (error) {
          // @ts-ignore
          // TODO: it's work, but it's not good, need to fix it
          setRoot([data]);
        }
      });
    // http://127.0.0.1:3000/api/v1/getusers
    fetch(address + "/api/v1/getusers", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: User[]) => {
        console.log(data);
        setUsers(data);
      });
  }
  useEffect(() => {
    Fetch();
  }, []);

  return (
    // @ts-ignore
    <ManageTasks>
      <Box>
        {root?.map((item) => (
          <>
            {/* @ts-ignore sx={(theme) => BoxTheme(theme)} */}
            <Box
              id={String(item.ID)}
              sx={(theme) => ({ padding: theme.spacing.lg })}
            >
              <Text
                sx={{
                  fontSize: "3em",
                  fontWeight: "bold",
                  borderBottom: "2px dashed",
                }}
              >
                {item.Name}
              </Text>
            </Box>
            {/* @ts-ignore */}
            <Box sx={(theme) => BoxTheme(theme)} id={String(item.ID)}>
              {item.Tasks.map((task) => (
                <div key={task.ID}>
                  <Text sx={{ fontSize: 25 }}>{task.Name}</Text>
                  <Text
                    sx={{
                      fontSize: 15,
                      textOverflow: "ellipsis",
                      overflow: "hidden",
                      maxWidth: "30vw",
                      whiteSpace: "nowrap",
                    }}
                  >
                    {task.Description}
                  </Text>
                </div>
              ))}
            </Box>
          </>
        ))}
      </Box>
    </ManageTasks>
  );
}
