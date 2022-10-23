# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgenden Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## 0) Mit dem Code vertraut machen

Öffnen Sie die Datei `frickelbude.go`. Und betrachten Sie den Code. Die Datei enthält folgende Definitionen:

1. Preisangaben
    - `PricingCPU`: Die monatlichen Kosten für die CPUs
    - `PricingRAM`: Die monatlichen Kosten für das Memory
    - `PricingSSD`: Die monatlichen Kosten für den Storage
2. Datentypen
    - `Server`: Die Struktur enthält Angaben für CPU, RAM und SSD und kann die Konfiguration sowohl für physische wie auch virtuelle Maschinen abbilden.
    - `GuestInventory`: Die Map speichert die Konfiguration für virtuelle Maschinen unter deren Namen ab.
    - `Host`: Die Struktur kombiniert die Hardware eines physischen Servers mit einer Reihe von Gast-VMs.
    - `HostInventory`: Die Map speichert die Konfiguration für physische Maschinen (inkl. Gast-Inventar) unter deren Namen ab.
3. Rechenzentrum
    - `DataCenter`: Dieses `HostInventory` ist der Hardware-Anfangsbestand im Rechenzentrum der FrickelCloud. Es gibt zu Beginn nur drei Server, auf denen je mehrere Gäste-VMs laufen.

Beantworten Sie die folgenden Fragen um Ihr Verständnis für diese Definitionen zu überprüfen:

1. Wie heissen die Gast-VMs auf dem Host `small-1`?
2. Wie viele CPU-Kerne beanspruchen die Gast-VMs auf dem Host `medium-1`? 
3. Über wie viel unzugeordneten SSD-Speicherplatz verfügt der Host `big-1`?

Schreiben Sie Ihre Antworten in die Datei `ex0/questions.md`.