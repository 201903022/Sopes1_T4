
# Sopes1_T4

## 📌 About

This test was performed on **Fedora 41**,  
using **local kubectl** and **Minikube** to create a local Kubernetes environment.

---

## 📥 Install kubectl locally

1. Open the following link in your browser:

```
https://dl.k8s.io/release/v1.29.2/bin/linux/amd64/kubectl
```

2. Navigate to the directory where the file was downloaded, then run:

```bash
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```

### ✅ Verify installation:

```bash
kubectl version --client
```

---

## 📦 Install Minikube locally

1. Download Minikube binary:

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
```

2. Make it executable and move it to your system path:

```bash
chmod +x minikube-linux-amd64
sudo mv minikube-linux-amd64 /usr/local/bin/minikube
```

3. Verify the installation:

```bash
minikube version
```

---

## 🚀 Start Minikube

If you have Docker installed, use the following to start Minikube:

```bash
minikube start --driver=docker
```

Otherwise, you can simply run:

```bash
minikube start
```

---

## ✅ Verify your Kubernetes cluster is running

```bash
kubectl get nodes
```

You should see something like:

```
NAME       STATUS   ROLES           AGE   VERSION
minikube   Ready    control-plane   ...   v1.32.0
```

---

## 🌐 Optional: Enable Ingress

If you plan to use domain-based routing (e.g., go.local):

```bash
minikube addons enable ingress
```

---

Feel free to customize this README for your project!