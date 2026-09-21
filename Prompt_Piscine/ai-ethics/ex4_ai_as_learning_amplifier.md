# Exercise 4: AI as Learning Amplifier

## Phase 1 – My attempt first (read Wikipedia + Cisco docs myself, no AI)

- Wireless routing: like OSPF/BGP but for WiFi/mesh.
- Protocols: AODV (on-demand), OLSR (proactive), Zigbee uses AODV-like for IoT.
- For small network (5-10 devices): maybe OLSR - low overhead, good for changing links.
- Justification: devices move, links break often in wireless.

## Phase 2 – Strategic questions to AI

- Asked: "In a mesh of 10 devices, why might OLSR be better than AODV? Edge cases?"

- Learned: OLSR floods topology - more reliable but higher battery use. AODV only when needed and saves power but slower setup.

## Phase 3 – Smart-city design (1000 sensors, 50 lights, 10 vehicles)

- Choice: Hybrid - LoRaWAN for long-range sensors (low power), WiFi 6 mesh for lights/vehicles (higher bandwidth).
- Failure points: Interference, battery death on sensors, jamming attacks.

- Refine with AI feedback: Add redundancy (multiple gateways), use MQTT over protocols for pub/sub.

### Reflection:
- About 70% my judgment (picked protocols from docs), 30% AI (edge cases, trade-offs).
- I could defend: "Low-power for sensors, mesh for real-time traffic."
- In 6 months: I'll remember power vs reliability trade-off.
- AI made me sharper this time. I asked targeted questions and got deeper than solo responses.