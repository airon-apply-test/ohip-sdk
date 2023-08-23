/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
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
