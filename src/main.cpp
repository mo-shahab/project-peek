#include <iostream>
#include <dirent.h>
#include <sys/stat.h>
#include <cstring>
#include <vector>

struct Node 
{
    std::string name;
    bool is_dir;
    std::vector<Node*> children;
};

void recursiveTraversal(const std::string& path) 
{
    DIR* dir = opendir(path.c_str());
    
    if (!dir) 
    {
        std::cerr << "Error: Cannot open directory " << path << "\n";
        return;
    }

    struct dirent* entry;
    
    while ((entry = readdir(dir)) != nullptr) 
    {
        std::string name = entry->d_name;

        // Skip current and parent directories
        if (name == "." || name == "..")
            continue;

        // Skip .git entirely
        if (name == ".git")
            continue;

        std::string fullPath = path + "/" + name;
        std::cout << fullPath << "\n";

        // Check if directory and recurse
        struct stat statbuf;
        if (stat(fullPath.c_str(), &statbuf) == 0 && S_ISDIR(statbuf.st_mode)) 
        {
            recursiveTraversal(fullPath);
        }
    }

    closedir(dir);
}

Node * build_tree(const std::string& path, const std::string& name)
{

    Node* node = new Node{name, true, {}};
    
    DIR* dir = opendir(path.c_str());
    if (!dir) 
    {
        std::cerr << "Error: Cannot open directory " << path << "\n";
        return node;
    }

    struct dirent* entry;
    while ((entry = readdir(dir)) != nullptr) 
    {
        std::string entry_name = entry->d_name;

        // Skip current and parent directories
        if (entry_name == "." || entry_name == "..")
            continue;

        // Skip .git entirely
        if (entry_name == ".git")
            continue;

        std::string fullpath = path + "/" + entry_name;
        std::cout << fullpath << "\n";

        // Check if directory and recurse
        struct stat statbuf;
        if (stat(fullpath.c_str(), &statbuf)) 
        {
            if(S_ISDIR(statbuf.st_mode))
            {
                Node* child = build_tree(fullpath, entry_name);
                node->children.push_back(child);
            }
            else 
            {
                Node* child = new Node{entry_name, false, {}};
                node->children.push_back(child);
            }
        }
    }

    closedir(dir);
    return node;
}

void print_tree(const Node* root, int depth = 0)
{
    std::string indent(depth* 4, ' ');
    std::cout << root->name << (root->is_dir ? "/" : "") << "\n";

    for(const auto& child: root->children)
    {
        print_tree(child, depth + 1);
    }
}

int main(int argc, char* argv[]) 
{
    if (argc < 2) 
    {
        std::cerr << "Usage: " << argv[0] << " <directory>\n";
        return 1;
    }

    std::string root_path = argv[1];

    Node* root = build_tree(root_path, root_path);
    print_tree(root, 0);

    return 0;
}

