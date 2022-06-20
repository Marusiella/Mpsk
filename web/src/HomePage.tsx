import React, { useEffect, useMemo, useState } from "react";
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
import { adress } from "./main";
import useSWR from "swr";

export interface Task {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt?: any;
  Name: string;
  Description: string;
  Done: boolean;
}

export interface RootObject {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt?: any;
  Name: string;
  Users?: any;
  Tasks: Task[];
}

// const fetcher = (...args) => fetch(...args).then((res) => res.json());

function BoxTheme(theme: MantineTheme) {
  return {
    backgroundColor:
      theme.colorScheme === "dark"
        ? theme.colors.dark[6]
        : theme.colors.gray[0],
    textAlign: "left",
    padding: theme.spacing.xl,
    borderRadius: theme.radius.md,
    cursor: "pointer",

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[5]
          : theme.colors.gray[1],
    },
  };
}

export default function HomePage() {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
  const [root, setRoot] = useState<RootObject[]>([]);
  function Fetch() {
    fetch(adress + "/api/v1/gettasks", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: RootObject[]) => {
        console.log(data);
        setRoot(data);
      });
  }
  useEffect(() => {
    Fetch();
  }, []);

  return (
    <AppShell
      styles={{
        main: {
          background:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
        },
      }}
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      fixed
      navbar={
        <Navbar
          p="md"
          hiddenBreakpoint="sm"
          hidden={!opened}
          width={{ sm: 200, lg: 300 }}
        >
          <Navbar.Section>{/* Header with logo */}</Navbar.Section>
          <Navbar.Section grow mt="md">
            {/* Links sections */}
          </Navbar.Section>
          <Navbar.Section>{/* Footer with user */}</Navbar.Section>
        </Navbar>
      }
      aside={
        <MediaQuery smallerThan="sm" styles={{ display: "none" }}>
          <Aside p="md" hiddenBreakpoint="sm" width={{ sm: 180, lg: 280 }}>
            <Text>Application sidebar</Text>
          </Aside>
        </MediaQuery>
      }
      header={
        <Header height={70} p="md">
          <div
            style={{ display: "flex", alignItems: "center", height: "100%" }}
          >
            <MediaQuery largerThan="sm" styles={{ display: "none" }}>
              <Burger
                opened={opened}
                onClick={() => setOpened((o) => !o)}
                size="sm"
                color={theme.colors.gray[6]}
                mr="xl"
              />
            </MediaQuery>

            <Text weight={"bold"} style={{ fontSize: "40px" }}>
              Mpsk
            </Text>
          </div>
        </Header>
      }
    >
      {root?.map((item) => (
        <>
          {/* @ts-ignore */}
          <Box sx={(theme) => BoxTheme(theme)} id={String(item.ID)}>
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
                <Text>{task.Name}</Text>
              </div>
            ))}
          </Box>
        </>
      ))}
    </AppShell>
  );
}
