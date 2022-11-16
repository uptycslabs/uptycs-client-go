package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		id      string
		out     interface{}
	}

	theTests := []convTest{
		{
			name:    "TestCustomProfile",
			fixture: "fixtures/customProfile.json",
			id:      "c6815103-33eb-41e0-bc2f-6a23cc2e1589",
			out: CustomProfile{
				ID:             "c6815103-33eb-41e0-bc2f-6a23cc2e1589",
				Name:           "custom ubuntu",
				Description:    "",
				QuerySchedules: CustomJSONString(`{"acpi_tables":3600,"apk_packages":7200,"apparmor_profiles":19800,"apt_sources":19800,"arp_cache":600,"atom_packages":43200,"audit_status":600,"authorized_keys":3600,"block_devices":3600,"carbon_black_info":7200,"certificates":3600,"chrome_extension_content_scripts":7200,"chrome_extensions":7200,"compliance":21600,"compliance_on_demand":1200,"cpu_time":300,"cpuid":3600,"crio_container_labels":600,"crio_container_mounts":600,"crio_container_stats":600,"crio_containers":600,"crio_image_fs_info":600,"crio_images":3600,"crio_pod_sandbox_labels":600,"crio_pod_sandboxes":600,"crio_status":3600,"crio_version":3600,"crontab":7200,"deb_packages":7200,"diag_watcher_stats":3600,"disk_encryption":3600,"dns_resolvers":600,"docker_container_labels":600,"docker_container_mounts":600,"docker_container_networks":600,"docker_container_ports":600,"docker_container_processes":30,"docker_containers":30,"docker_image_history":3600,"docker_image_labels":3600,"docker_image_layers":3600,"docker_images":3600,"docker_info":21600,"docker_network_labels":3600,"docker_networks":3600,"docker_version":21600,"docker_volume_labels":3600,"docker_volumes":3600,"ebpf_kernel_support":3600,"ec2_instance_metadata":7200,"efivars":21600,"etc_hosts":19800,"etc_protocols":19800,"etc_services":19800,"firefox_addons":7200,"groups":3600,"interface_addresses":3600,"interface_details":600,"interface_ipv6":21600,"iptables":19800,"kernel_info":3600,"kernel_integrity":1800,"kernel_modules":1200,"known_hosts":3600,"last":3600,"listening_ports":19800,"load_average":60,"logged_in_users":60,"lxd_certificates":3600,"lxd_cluster":600,"lxd_cluster_members":600,"lxd_images":3600,"lxd_instances":30,"lxd_networks":3600,"lxd_storage_pools":3600,"md_devices":21600,"md_drives":21600,"md_personalities":21600,"memory_array_mapped_addresses":19800,"memory_arrays":19800,"memory_device_mapped_addresses":19800,"memory_devices":21600,"memory_error_info":19800,"memory_info":300,"memory_map":3600,"mounts":600,"msr":3600,"npm_packages":21600,"oem_strings":21600,"opera_extensions":7200,"os_version":7200,"osquery_config":3600,"osquery_events":600,"osquery_extensions":3600,"osquery_flags":3600,"osquery_info":3600,"osquery_packs":3600,"osquery_registry":3600,"osquery_schedule":19800,"osquery_upt_stats":600,"pci_devices":3600,"platform_info":7200,"process_cpu":300,"process_envs":300,"process_namespaces":21600,"process_open_files":300,"process_open_pipes":600,"process_open_sockets_local":30,"process_open_sockets_remote":30,"processes":30,"processes_hash":300,"python_packages":19800,"routes":3600,"rpm_packages":7200,"selinux_settings":19800,"shadow":300,"shared_memory":3600,"slack_user_info":3600,"smbios_tables":3600,"ssh_configs":21600,"startup_items":7200,"sudoers":3600,"suid_bin":3600,"system_controls":19800,"system_info":21600,"ulimit_info":21600,"upt_op_interfaces":7200,"uptime":3600,"usb_devices":3600,"user_groups":3600,"user_ssh_keys":3600,"users":300,"vulnerabilities":21600}`),
				Priority:       1,
				CreatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:      "2021-06-15T21:14:22.001Z",
				UpdatedAt:      "2021-06-15T21:44:30.182Z",
				ResourceType:   "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Custom profile information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles/c6815103-33eb-41e0-bc2f-6a23cc2e1589"},
					LinkItem{Rel: "parent", Title: "Custom profiles information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/customProfiles/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			customProfileResp, err := c.GetCustomProfile(CustomProfile{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(customProfileResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", customProfileResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteCustomProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   CustomProfile
	}

	theTests := []convTest{
		{
			name: "TestCustomProfile",
			in: CustomProfile{
				ID:             "9cde7195-ec0c-475e-a208-dbf81a32798a",
				Name:           "custom ubuntu",
				Description:    "",
				QuerySchedules: CustomJSONString(`{"acpi_tables":3600,"apk_packages":7200,"apparmor_profiles":19800,"apt_sources":19800,"arp_cache":600,"atom_packages":43200,"audit_status":600,"authorized_keys":3600,"block_devices":3600,"carbon_black_info":7200,"certificates":3600,"chrome_extension_content_scripts":7200,"chrome_extensions":7200,"compliance":21600,"compliance_on_demand":1200,"cpu_time":300,"cpuid":3600,"crio_container_labels":600,"crio_container_mounts":600,"crio_container_stats":600,"crio_containers":600,"crio_image_fs_info":600,"crio_images":3600,"crio_pod_sandbox_labels":600,"crio_pod_sandboxes":600,"crio_status":3600,"crio_version":3600,"crontab":7200,"deb_packages":7200,"diag_watcher_stats":3600,"disk_encryption":3600,"dns_resolvers":600,"docker_container_labels":600,"docker_container_mounts":600,"docker_container_networks":600,"docker_container_ports":600,"docker_container_processes":30,"docker_containers":30,"docker_image_history":3600,"docker_image_labels":3600,"docker_image_layers":3600,"docker_images":3600,"docker_info":21600,"docker_network_labels":3600,"docker_networks":3600,"docker_version":21600,"docker_volume_labels":3600,"docker_volumes":3600,"ebpf_kernel_support":3600,"ec2_instance_metadata":7200,"efivars":21600,"etc_hosts":19800,"etc_protocols":19800,"etc_services":19800,"firefox_addons":7200,"groups":3600,"interface_addresses":3600,"interface_details":600,"interface_ipv6":21600,"iptables":19800,"kernel_info":3600,"kernel_integrity":1800,"kernel_modules":1200,"known_hosts":3600,"last":3600,"listening_ports":19800,"load_average":60,"logged_in_users":60,"lxd_certificates":3600,"lxd_cluster":600,"lxd_cluster_members":600,"lxd_images":3600,"lxd_instances":30,"lxd_networks":3600,"lxd_storage_pools":3600,"md_devices":21600,"md_drives":21600,"md_personalities":21600,"memory_array_mapped_addresses":19800,"memory_arrays":19800,"memory_device_mapped_addresses":19800,"memory_devices":21600,"memory_error_info":19800,"memory_info":300,"memory_map":3600,"mounts":600,"msr":3600,"npm_packages":21600,"oem_strings":21600,"opera_extensions":7200,"os_version":7200,"osquery_config":3600,"osquery_events":600,"osquery_extensions":3600,"osquery_flags":3600,"osquery_info":3600,"osquery_packs":3600,"osquery_registry":3600,"osquery_schedule":19800,"osquery_upt_stats":600,"pci_devices":3600,"platform_info":7200,"process_cpu":300,"process_envs":300,"process_namespaces":21600,"process_open_files":300,"process_open_pipes":600,"process_open_sockets_local":30,"process_open_sockets_remote":30,"processes":30,"processes_hash":300,"python_packages":19800,"routes":3600,"rpm_packages":7200,"selinux_settings":19800,"shadow":300,"shared_memory":3600,"slack_user_info":3600,"smbios_tables":3600,"ssh_configs":21600,"startup_items":7200,"sudoers":3600,"suid_bin":3600,"system_controls":19800,"system_info":21600,"ulimit_info":21600,"upt_op_interfaces":7200,"uptime":3600,"usb_devices":3600,"user_groups":3600,"user_ssh_keys":3600,"users":300,"vulnerabilities":21600}`),
				Priority:       1,
				CreatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:      "2021-06-15T21:14:22.001Z",
				UpdatedAt:      "2021-06-15T21:44:30.182Z",
				ResourceType:   "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Custom profile information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles/c6815103-33eb-41e0-bc2f-6a23cc2e1589"},
					LinkItem{Rel: "parent", Title: "Custom profiles information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/customProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteCustomProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/customProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutCustomProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      CustomProfile
	}

	theTests := []convTest{
		{
			name:    "TestCustomProfile",
			fixture: "fixtures/customProfile.json",
			in: CustomProfile{
				ID:             "c6815103-33eb-41e0-bc2f-6a23cc2e1589",
				Name:           "custom ubuntu",
				Description:    "",
				QuerySchedules: CustomJSONString(`{"acpi_tables":3600,"apk_packages":7200,"apparmor_profiles":19800,"apt_sources":19800,"arp_cache":600,"atom_packages":43200,"audit_status":600,"authorized_keys":3600,"block_devices":3600,"carbon_black_info":7200,"certificates":3600,"chrome_extension_content_scripts":7200,"chrome_extensions":7200,"compliance":21600,"compliance_on_demand":1200,"cpu_time":300,"cpuid":3600,"crio_container_labels":600,"crio_container_mounts":600,"crio_container_stats":600,"crio_containers":600,"crio_image_fs_info":600,"crio_images":3600,"crio_pod_sandbox_labels":600,"crio_pod_sandboxes":600,"crio_status":3600,"crio_version":3600,"crontab":7200,"deb_packages":7200,"diag_watcher_stats":3600,"disk_encryption":3600,"dns_resolvers":600,"docker_container_labels":600,"docker_container_mounts":600,"docker_container_networks":600,"docker_container_ports":600,"docker_container_processes":30,"docker_containers":30,"docker_image_history":3600,"docker_image_labels":3600,"docker_image_layers":3600,"docker_images":3600,"docker_info":21600,"docker_network_labels":3600,"docker_networks":3600,"docker_version":21600,"docker_volume_labels":3600,"docker_volumes":3600,"ebpf_kernel_support":3600,"ec2_instance_metadata":7200,"efivars":21600,"etc_hosts":19800,"etc_protocols":19800,"etc_services":19800,"firefox_addons":7200,"groups":3600,"interface_addresses":3600,"interface_details":600,"interface_ipv6":21600,"iptables":19800,"kernel_info":3600,"kernel_integrity":1800,"kernel_modules":1200,"known_hosts":3600,"last":3600,"listening_ports":19800,"load_average":60,"logged_in_users":60,"lxd_certificates":3600,"lxd_cluster":600,"lxd_cluster_members":600,"lxd_images":3600,"lxd_instances":30,"lxd_networks":3600,"lxd_storage_pools":3600,"md_devices":21600,"md_drives":21600,"md_personalities":21600,"memory_array_mapped_addresses":19800,"memory_arrays":19800,"memory_device_mapped_addresses":19800,"memory_devices":21600,"memory_error_info":19800,"memory_info":300,"memory_map":3600,"mounts":600,"msr":3600,"npm_packages":21600,"oem_strings":21600,"opera_extensions":7200,"os_version":7200,"osquery_config":3600,"osquery_events":600,"osquery_extensions":3600,"osquery_flags":3600,"osquery_info":3600,"osquery_packs":3600,"osquery_registry":3600,"osquery_schedule":19800,"osquery_upt_stats":600,"pci_devices":3600,"platform_info":7200,"process_cpu":300,"process_envs":300,"process_namespaces":21600,"process_open_files":300,"process_open_pipes":600,"process_open_sockets_local":30,"process_open_sockets_remote":30,"processes":30,"processes_hash":300,"python_packages":19800,"routes":3600,"rpm_packages":7200,"selinux_settings":19800,"shadow":300,"shared_memory":3600,"slack_user_info":3600,"smbios_tables":3600,"ssh_configs":21600,"startup_items":7200,"sudoers":3600,"suid_bin":3600,"system_controls":19800,"system_info":21600,"ulimit_info":21600,"upt_op_interfaces":7200,"uptime":3600,"usb_devices":3600,"user_groups":3600,"user_ssh_keys":3600,"users":300,"vulnerabilities":21600}`),
				Priority:       1,
				CreatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:      "2021-06-15T21:14:22.001Z",
				UpdatedAt:      "2021-06-15T21:44:30.182Z",
				ResourceType:   "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Custom profile information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles/c6815103-33eb-41e0-bc2f-6a23cc2e1589"},
					LinkItem{Rel: "parent", Title: "Custom profiles information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/customProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateCustomProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/customProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateCustomProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      CustomProfile
	}

	theTests := []convTest{
		{
			name:    "TestCustomProfile",
			fixture: "fixtures/customProfile.json",
			in: CustomProfile{
				ID:             "c6815103-33eb-41e0-bc2f-6a23cc2e1589",
				Name:           "custom ubuntu",
				Description:    "",
				QuerySchedules: CustomJSONString(`{"acpi_tables":3600,"apk_packages":7200,"apparmor_profiles":19800,"apt_sources":19800,"arp_cache":600,"atom_packages":43200,"audit_status":600,"authorized_keys":3600,"block_devices":3600,"carbon_black_info":7200,"certificates":3600,"chrome_extension_content_scripts":7200,"chrome_extensions":7200,"compliance":21600,"compliance_on_demand":1200,"cpu_time":300,"cpuid":3600,"crio_container_labels":600,"crio_container_mounts":600,"crio_container_stats":600,"crio_containers":600,"crio_image_fs_info":600,"crio_images":3600,"crio_pod_sandbox_labels":600,"crio_pod_sandboxes":600,"crio_status":3600,"crio_version":3600,"crontab":7200,"deb_packages":7200,"diag_watcher_stats":3600,"disk_encryption":3600,"dns_resolvers":600,"docker_container_labels":600,"docker_container_mounts":600,"docker_container_networks":600,"docker_container_ports":600,"docker_container_processes":30,"docker_containers":30,"docker_image_history":3600,"docker_image_labels":3600,"docker_image_layers":3600,"docker_images":3600,"docker_info":21600,"docker_network_labels":3600,"docker_networks":3600,"docker_version":21600,"docker_volume_labels":3600,"docker_volumes":3600,"ebpf_kernel_support":3600,"ec2_instance_metadata":7200,"efivars":21600,"etc_hosts":19800,"etc_protocols":19800,"etc_services":19800,"firefox_addons":7200,"groups":3600,"interface_addresses":3600,"interface_details":600,"interface_ipv6":21600,"iptables":19800,"kernel_info":3600,"kernel_integrity":1800,"kernel_modules":1200,"known_hosts":3600,"last":3600,"listening_ports":19800,"load_average":60,"logged_in_users":60,"lxd_certificates":3600,"lxd_cluster":600,"lxd_cluster_members":600,"lxd_images":3600,"lxd_instances":30,"lxd_networks":3600,"lxd_storage_pools":3600,"md_devices":21600,"md_drives":21600,"md_personalities":21600,"memory_array_mapped_addresses":19800,"memory_arrays":19800,"memory_device_mapped_addresses":19800,"memory_devices":21600,"memory_error_info":19800,"memory_info":300,"memory_map":3600,"mounts":600,"msr":3600,"npm_packages":21600,"oem_strings":21600,"opera_extensions":7200,"os_version":7200,"osquery_config":3600,"osquery_events":600,"osquery_extensions":3600,"osquery_flags":3600,"osquery_info":3600,"osquery_packs":3600,"osquery_registry":3600,"osquery_schedule":19800,"osquery_upt_stats":600,"pci_devices":3600,"platform_info":7200,"process_cpu":300,"process_envs":300,"process_namespaces":21600,"process_open_files":300,"process_open_pipes":600,"process_open_sockets_local":30,"process_open_sockets_remote":30,"processes":30,"processes_hash":300,"python_packages":19800,"routes":3600,"rpm_packages":7200,"selinux_settings":19800,"shadow":300,"shared_memory":3600,"slack_user_info":3600,"smbios_tables":3600,"ssh_configs":21600,"startup_items":7200,"sudoers":3600,"suid_bin":3600,"system_controls":19800,"system_info":21600,"ulimit_info":21600,"upt_op_interfaces":7200,"uptime":3600,"usb_devices":3600,"user_groups":3600,"user_ssh_keys":3600,"users":300,"vulnerabilities":21600}`),
				Priority:       1,
				CreatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy:      "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				CreatedAt:      "2021-06-15T21:14:22.001Z",
				UpdatedAt:      "2021-06-15T21:44:30.182Z",
				ResourceType:   "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Custom profile information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles/c6815103-33eb-41e0-bc2f-6a23cc2e1589"},
					LinkItem{Rel: "parent", Title: "Custom profiles information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/customProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/customProfiles",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateCustomProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/customProfiles"], 1)
		})
	}
}
