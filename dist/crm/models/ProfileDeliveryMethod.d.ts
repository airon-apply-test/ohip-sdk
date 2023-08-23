/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ProfileDeliveryModuleType } from './ProfileDeliveryModuleType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Delivery Information type to the profile.
 * @export
 * @interface ProfileDeliveryMethod
 */
export interface ProfileDeliveryMethod {
    /**
     *
     * @type {UniqueIDType}
     * @memberof ProfileDeliveryMethod
     */
    deliveryId?: UniqueIDType;
    /**
     * Delivery type can have a value EMAIL, ELECTRONIC etc and it depends on the parameter set in OPERA Control.
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    deliveryType?: string;
    /**
     * Delivery value holds the corresponding value of the delivery type..
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    deliveryValue?: string;
    /**
     * Property that has delivery methods configured.
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    hotelId?: string;
    /**
     *
     * @type {ProfileDeliveryModuleType}
     * @memberof ProfileDeliveryMethod
     */
    deliveryModule?: ProfileDeliveryModuleType;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof ProfileDeliveryMethod
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ProfileDeliveryMethod
     */
    orderSequence?: number;
}
/**
 * Check if a given object implements the ProfileDeliveryMethod interface.
 */
export declare function instanceOfProfileDeliveryMethod(value: object): boolean;
export declare function ProfileDeliveryMethodFromJSON(json: any): ProfileDeliveryMethod;
export declare function ProfileDeliveryMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileDeliveryMethod;
export declare function ProfileDeliveryMethodToJSON(value?: ProfileDeliveryMethod | null): any;
