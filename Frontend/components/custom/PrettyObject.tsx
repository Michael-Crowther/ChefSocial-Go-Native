import { Text, View } from "react-native";

export function PrettyObject({ children }: { children: any }) {
  return (
    <View className="border border-success">
      <Text className="border-b text-center">Object</Text>
      <Text>{JSON.stringify(children, null, "\t")}</Text>
    </View>
  );
}
