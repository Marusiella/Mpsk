import { Box } from "@mantine/core";
import React from "react";
import { address } from "./main";
import MainLayout from "./MainLayout";
import { Group } from "./models";

export default function ManageTasks() {
  const [root, setRoot] = React.useState<Group[]>([]);
  function Fetch() {
    fetch(address + "/api/v1/gettasks", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: Group[]) => {
        // console.log([data]);
        try {
          //   data.map(() => {});
          setRoot(data);
        } catch (error) {}
      });
  }
  return (
    <MainLayout>
      <Box>
        {root?.map((item) => (
          <>
            <Box id={String(item.ID)}>{item.Name}</Box>
          </>
        ))}
      </Box>
    </MainLayout>
  );
}
