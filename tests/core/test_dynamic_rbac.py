import time

from .common import random_str


def test_admin_can_see_links(admin_mc, remove_resource):
    cluster = admin_mc.client.create_cluster(name=random_str())
    remove_resource(cluster)

    user_cluster_view = admin_mc.client.by_id_cluster(cluster.id)

    assert 'update' in user_cluster_view.links
    assert 'remove' in user_cluster_view.links


def test_read_only_cannot_see_links(admin_mc, user_mc, remove_resource):
    cluster = admin_mc.client.create_cluster(name=random_str())
    remove_resource(cluster)

    time.sleep(5)  # TODO poll

    crtb = admin_mc.client.create_cluster_role_template_binding(
        clusterId=cluster.id,
        roleTemplateId="read-only",
        userId=user_mc.user.id)
    remove_resource(crtb)

    time.sleep(5)  # TODO poll

    user_cluster_view = user_mc.client.by_id_cluster(cluster.id)

    assert 'update' not in user_cluster_view.links
    assert 'remove' not in user_cluster_view.links
