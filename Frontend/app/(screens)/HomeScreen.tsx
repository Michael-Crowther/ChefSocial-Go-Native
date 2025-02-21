import React from "react";
import { View, Text, ScrollView } from "react-native";
import { useQuery } from "@tanstack/react-query";
import { PrettyObject } from "@/components/custom/PrettyObject";

async function fetchRecipes() {
  const response = await fetch("http://10.0.0.201:3005/recipes");
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json();
}

export default function HomeScreen() {
  const { data: recipes } = useQuery({
    queryKey: ["recipes"],
    queryFn: fetchRecipes,
  });

  return (
    <ScrollView className="p-6">
      <PrettyObject>{recipes}</PrettyObject>
    </ScrollView>
  );
}
