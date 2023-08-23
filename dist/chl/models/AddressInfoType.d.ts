/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AddressType } from './AddressType';
/**
 * Provides address information.
 * @export
 * @interface AddressInfoType
 */
export interface AddressInfoType {
    /**
     *
     * @type {AddressType}
     * @memberof AddressInfoType
     */
    address?: AddressType;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof AddressInfoType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof AddressInfoType
     */
    type?: string;
}
/**
 * Check if a given object implements the AddressInfoType interface.
 */
export declare function instanceOfAddressInfoType(value: object): boolean;
export declare function AddressInfoTypeFromJSON(json: any): AddressInfoType;
export declare function AddressInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddressInfoType;
export declare function AddressInfoTypeToJSON(value?: AddressInfoType | null): any;
