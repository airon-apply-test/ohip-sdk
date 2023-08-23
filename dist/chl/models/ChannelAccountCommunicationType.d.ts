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
import type { ChannelAccountCommunicationTypeAddresses } from './ChannelAccountCommunicationTypeAddresses';
/**
 * Holds channel account communication information.
 * @export
 * @interface ChannelAccountCommunicationType
 */
export interface ChannelAccountCommunicationType {
    /**
     *
     * @type {ChannelAccountCommunicationTypeAddresses}
     * @memberof ChannelAccountCommunicationType
     */
    addresses?: ChannelAccountCommunicationTypeAddresses;
}
/**
 * Check if a given object implements the ChannelAccountCommunicationType interface.
 */
export declare function instanceOfChannelAccountCommunicationType(value: object): boolean;
export declare function ChannelAccountCommunicationTypeFromJSON(json: any): ChannelAccountCommunicationType;
export declare function ChannelAccountCommunicationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelAccountCommunicationType;
export declare function ChannelAccountCommunicationTypeToJSON(value?: ChannelAccountCommunicationType | null): any;
